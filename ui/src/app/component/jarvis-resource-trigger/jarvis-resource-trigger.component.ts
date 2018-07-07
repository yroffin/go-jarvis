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

declare var Prism: any;

import { LoggerService } from '../../service/logger.service';
import { JarvisConfigurationService } from '../../service/jarvis-configuration.service';
import { JarvisResourceLink } from '../../class/jarvis-resource-link';

import { JarvisDataTriggerService } from '../../service/jarvis-data-trigger.service';
import { JarvisDataCronService } from '../../service/jarvis-data-cron.service';

/**
 * class
 */
import { JarvisResource } from '../../class/jarvis-resource';
import { NotifyCallback } from '../../class/jarvis-resource';

/**
 * data model
 */
import { ResourceBean } from '../../model/resource-bean';
import { TriggerBean } from '../../model/trigger-bean';
import { CronBean } from '../../model/cron-bean';
import { PickerBean } from '../../model/picker-bean';
import { MatTableDataSource } from '@angular/material';
import { JarvisPickResourceService } from '../../service/jarvis-pick-resource.service';
import { Observable } from 'rxjs';
import { ResourceStoreService } from '../../store/resources.store';

@Component({
  selector: 'app-jarvis-resource-trigger',
  templateUrl: './jarvis-resource-trigger.component.html',
  styleUrls: ['./jarvis-resource-trigger.component.css']
})
export class JarvisResourceTriggerComponent extends JarvisResource<TriggerBean> implements NotifyCallback<ResourceBean>, OnInit {

  myStream: Observable<TriggerBean>;
  @Input() myTrigger: TriggerBean;
  public myMatResources = new MatTableDataSource([]);
  public myMatDevicesResources = new MatTableDataSource([]);
  public myMatProcessesResources = new MatTableDataSource([]);

  /**
   * internal vars
   */
  myCron: CronBean;

  private jarvisCronLink: JarvisResourceLink<CronBean>;

  /**
   * constructor
   */
  constructor(
    private _route: ActivatedRoute,
    private _router: Router,
    private _jarvisConfigurationService: JarvisConfigurationService,
    private _triggerService: JarvisDataTriggerService,
    private logger: LoggerService,
    private _cronService: JarvisDataCronService,
    private resourceStoreService: ResourceStoreService,
    private jarvisPickResourceService: JarvisPickResourceService,
   ) {
    super('/triggers', ['execute'], _triggerService, _route, _router);
    this.jarvisCronLink = new JarvisResourceLink<CronBean>(this.logger);
    this.myStream = this.resourceStoreService.trigger();
  }

  /**
   * load device and related data
   */
  ngOnInit() {
    this.myStream.subscribe(
      (resource: TriggerBean) => {
        this.setResource(resource);
        let picker: PickerBean = new PickerBean();
        picker.action = 'complete';
        this.notify(picker, resource);
      }
    )
  }

  /**
   * highlight source
   */
  public hightlight(body: string): void {
    if(body) {
      return Prism.highlight(body, Prism.languages.clike);
    } else {
      return Prism.highlight("// empty", Prism.languages.clike);
    }
  }

  /**
   * complete resource
   */
  public complete(resource: TriggerBean): void {
    this.myTrigger = resource;
    this.myTrigger.crons = [];
    (new JarvisResourceLink<CronBean>(this.logger)).loadLinksWithCallback(resource.id, resource.crons, this._triggerService.allLinkedCron, (elements) => {
        this.myTrigger.crons = elements;
      });
  }

  /**
   * task action
   */
  public execute(): void {
    /**
     * execute this plugin
     */
    let myOutputData;
    this._triggerService.Task(this.myTrigger.id, 'execute', {})
      .subscribe(
      (result: any) => myOutputData = result,
      error => console.log(error),
      () => {
      }
      );
  }

    /**
   * task action
   */
  public collect(): void {
    let output: any;
    this._triggerService.Task(this.myTrigger.id, 'collect', {})
      .subscribe(
      (result: any) => {
      },
      error => console.log(error),
      () => {
      }
      );
  }

  /**
   * pick action
   */
  public pick(picker: PickerBean): void {
    /**
     * find notifications
     */
    if (picker.action === 'crons') {
      this.jarvisPickResourceService.open(this, 'Crontabs', picker);
    }
  }

  /**
   * notify to add new resource
   */
  public notify(picker: PickerBean, resource: ResourceBean): void {
    if (picker.action === 'crons') {
      this.jarvisCronLink.addLink(this.getResource().id, resource.id, this.getResource().crons, { "order": "1", href: "HREF" }, this._triggerService.allLinkedCron);
    }
    if (picker.action === 'complete') {
      this.myTrigger = <TriggerBean>resource;
      this.myTrigger.crons = [];
      (new JarvisResourceLink<CronBean>(this.logger)).loadLinksWithCallback(resource.id, this.myTrigger.crons, this._triggerService.allLinkedCron, (elements) => {
        this.myTrigger.crons = elements;
        this.myMatResources = new MatTableDataSource(elements);
        this.myMatDevicesResources = new MatTableDataSource(this.myTrigger.devices);
        this.myMatProcessesResources = new MatTableDataSource(this.myTrigger.processes);
      });
    }
  }

  /**
   * drop crontab link
   */
  public dropCronLink(linked: CronBean): void {
    this.jarvisCronLink.dropLink(linked, this.myTrigger.id, this.myTrigger.crons, this._triggerService.allLinkedCron);
  }

  /**
   * drop crontab link
   */
  public updateCronLink(linked: CronBean): void {
    this.jarvisCronLink.updateLink(linked, this.myTrigger.id, this._triggerService.allLinkedCron);
  }

  /**
   * goto crontab link
   */
  public gotoCronLink(linked: CronBean): void {
    this._router.navigate(['/crons/' + linked.id]);
  }
}
