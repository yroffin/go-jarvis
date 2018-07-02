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

import { JarvisConfigurationService } from '../../service/jarvis-configuration.service';
import { JarvisResourceLink } from '../../class/jarvis-resource-link';

import { JarvisDataConfigurationService } from '../../service/jarvis-data-configuration.service';

/**
 * class
 */
import { JarvisResource } from '../../class/jarvis-resource';
import { NotifyCallback } from '../../class/jarvis-resource';

/**
 * data model
 */
import { ResourceBean } from '../../model/resource-bean';
import { PickerBean } from '../../model/picker-bean';
import { ConfigurationBean } from '../../model/configuration-bean';
import { MatTableDataSource } from '@angular/material';
import { Observable } from 'rxjs';
import { ResourceStoreService } from '../../store/resources.store';

@Component({
  selector: 'app-jarvis-resource-configuration',
  templateUrl: './jarvis-resource-configuration.component.html',
  styleUrls: ['./jarvis-resource-configuration.component.css']
})
export class JarvisResourceConfigurationComponent extends JarvisResource<ConfigurationBean> implements NotifyCallback<ResourceBean>, OnInit {

  myStream: Observable<ConfigurationBean>;
  @Input() myConfiguration: ConfigurationBean;
  public myMatResources = new MatTableDataSource([]);

  /**
   * constructor
   */
  constructor(
    private _route: ActivatedRoute,
    private _router: Router,
    private _jarvisConfigurationService: JarvisConfigurationService,
    private resourceStoreService: ResourceStoreService,
    private _configurationService: JarvisDataConfigurationService
  ) {
    super('/configurations', [], _configurationService, _route, _router);
    this.myStream = this.resourceStoreService.configuration();
  }

  /**
   * load device and related data
   */
  ngOnInit() {
    this.myStream.subscribe(
      (resource: ConfigurationBean) => {
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
   * notify to add new resource
   */
  public notify(picker: PickerBean, resource: ResourceBean): void {
    if (picker.action === 'complete') {
      this.myConfiguration = resource;
      this.myMatResources = new MatTableDataSource(this.myConfiguration.system.scheduled)
    }
  }
}
