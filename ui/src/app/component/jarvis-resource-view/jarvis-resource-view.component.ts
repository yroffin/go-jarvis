/* 
 * Copyright 2016 Yannick Roffin.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, Input, ViewChild, OnInit, ElementRef } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { SelectItem } from 'primeng/primeng';
import { ConfirmationService } from 'primeng/primeng';

import { DataSource } from '@angular/cdk/collections';
import { Observable, BehaviorSubject } from 'rxjs';

declare var Prism: any;

import { LoggerService } from '../../service/logger.service';
import { JarvisConfigurationService } from '../../service/jarvis-configuration.service';
import { JarvisDataDeviceService } from '../../service/jarvis-data-device.service';
import { JarvisDataViewService } from '../../service/jarvis-data-view.service';
import { JarvisResourceLink } from '../../class/jarvis-resource-link';

/**
 * class
 */
import { JarvisResource } from '../../class/jarvis-resource';
import { NotifyCallback } from '../../class/jarvis-resource';
import { MatTableDataSource } from '@angular/material';

/**
 * data model
 */
import { ResourceBean } from '../../model/resource-bean';
import { PickerBean } from '../../model/picker-bean';
import { ViewBean } from '../../model/view-bean';
import { DeviceBean } from '../../model/device-bean';
import { JarvisPickResourceService } from '../../service/jarvis-pick-resource.service';
import { JarvisDataConnectorService } from '../../service/jarvis-data-connector.service.';
import { ResourceStoreService } from '../../store/resources.store';

@Component({
  selector: 'app-jarvis-resource-view',
  templateUrl: './jarvis-resource-view.component.html',
  styleUrls: ['./jarvis-resource-view.component.css']
})
export class JarvisResourceViewComponent extends JarvisResource<ViewBean> implements NotifyCallback<ResourceBean>, OnInit {

  myStream: Observable<ViewBean>;
  @Input() myView: ViewBean;
  public devices = new MatTableDataSource([])

  /**
   * internal vars
   */
  private jarvisDeviceLink: JarvisResourceLink<DeviceBean>;

  /**
   * constructor
   */
  constructor(
    private jarvisPickResourceService: JarvisPickResourceService,
    private _route: ActivatedRoute,
    private _router: Router,
    private _jarvisConfigurationService: JarvisConfigurationService,
    private _viewService: JarvisDataViewService,
    private logger: LoggerService,    
    private resourceStoreService: ResourceStoreService,
    private _deviceService: JarvisDataDeviceService) {
    super('/views', [], _viewService, _route, _router);
    this.jarvisDeviceLink = new JarvisResourceLink<DeviceBean>(this.logger);
    this.myStream = resourceStoreService.view();
  }

  /**
   * load device and related data
   */
  ngOnInit() {
    this.myStream.subscribe(
      (resource: ViewBean) => {
        this.setResource(resource);
        let picker: PickerBean = new PickerBean();
        picker.action = 'complete';
        this.notify(picker, resource);
      }
    )
  }

  /**
   * task action
   */
  public task(action: string): void {
  }

  /**
   * pick action
   */
  public pick(picker: PickerBean): void {
    /**
     * find notifications
     */
    if (picker.action === 'devices') {
      this.jarvisPickResourceService.open(this, 'Device', picker);
    }
  }

  /**
   * notify to add new resource
   */
  public notify(picker: PickerBean, resource: ResourceBean): void {
    if( picker.action === 'devices') {
      this.jarvisDeviceLink.addLink(this.getResource().id, resource.id, this.getResource().devices, {"order": "1", href: "HREF"}, this._viewService.allLinkedDevice);
    }
    if( picker.action === 'complete') {
      this.myView = <ViewBean> resource;
      this.myView.devices = [];
      (new JarvisResourceLink<DeviceBean>(this.logger)).loadLinksWithCallback(resource.id, this.myView.devices, this._viewService.allLinkedDevice, (elements) => {
        this.myView.devices = elements;
        this.devices = new MatTableDataSource(elements)
      });
    }
  }

  /**
   * drop link
   */
  public dropDeviceLink(linked: DeviceBean): void {
    this.jarvisDeviceLink.dropLink(linked, this.myView.id, this.myView.devices, this._viewService.allLinkedDevice);
  }

  /**
   * drop link
   */
  public updateDeviceLink(linked: DeviceBean): void {
    this.jarvisDeviceLink.updateLink(linked, this.myView.id, this._viewService.allLinkedDevice);
  }

  /**
   * goto link
   */
  public gotoDeviceLink(linked: DeviceBean): void {
    this._router.navigate(['/devices/' + linked.id]);
  }
}
