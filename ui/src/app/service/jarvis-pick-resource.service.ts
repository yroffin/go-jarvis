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

import { Injectable, Inject } from '@angular/core';
import { ResourceBean } from '../model/resource-bean';
import { JarvisDataDeviceService } from './jarvis-data-device.service';
import { JarvisDataTriggerService } from './jarvis-data-trigger.service';
import { JarvisDataPluginService } from './jarvis-data-plugin.service';
import { JarvisDataCommandService } from './jarvis-data-command.service';
import { JarvisDataProcessService } from './jarvis-data-process.service';
import { JarvisDataCronService } from './jarvis-data-cron.service';
import { JarvisDataDatasourceService } from './jarvis-data-datasource.service';
import { JarvisDataConnectorService } from './jarvis-data-connector.service.';
import { JarvisDataNotificationService } from './jarvis-data-notification.service';
import { LoggerService } from './logger.service';
import { JarvisDefaultResource } from '../interface/jarvis-default-resource';
import { PickerBean } from '../model/picker-bean';
import { NotifyCallback } from '../class/jarvis-resource';
import { DialogPickResource } from '../dialog/picker/jarvis-pick-resource.component';
import { MatDialog } from '@angular/material';

@Injectable({
  providedIn: 'root'
})
export class JarvisPickResourceService {

  /**
   * constructor
   */
  constructor(
    private connectorService: JarvisDataConnectorService,
    private deviceService: JarvisDataDeviceService,
    private triggerService: JarvisDataTriggerService,
    private pluginService: JarvisDataPluginService,
    private commandService: JarvisDataCommandService,
    private processService: JarvisDataProcessService,
    private cronService: JarvisDataCronService,
    private datasourceService: JarvisDataDatasourceService,
    private notificationService: JarvisDataNotificationService,
    private logger: LoggerService,
    private dialog: MatDialog
  ) {
  }

  /**
   * open dialog resource
   * @param target 
   * @param name 
   * @param resource 
   */
  public open(target: NotifyCallback<ResourceBean>, name: string, picker: PickerBean) {
    // Get service and load data
    this.getService(picker.action).GetAll()
      .subscribe(
      (data: ResourceBean[]) => {
        const dialogRef = this.dialog.open(DialogPickResource, {
          width: '800px',
          data: {
            target: target,
            title: name,
            resources: data,
            picker: picker,
          }
        });

        dialogRef.afterClosed().subscribe(result => {
        });
      },
      (error) => {
        this.logger.error("In loadResource", error);
      });
  }

  private getService(name: string): JarvisDefaultResource<ResourceBean> {
    let service: JarvisDefaultResource<ResourceBean>;
    if (name === 'connectors') {
      service = this.connectorService;
    }
    if (name === 'processes') {
      service = this.processService;
    }
    if (name === 'crons') {
      service = this.cronService;
    }
    if (name === 'devices') {
      service = this.deviceService;
    }
    if (name === 'triggers') {
      service = this.triggerService;
    }
    if (name === 'plugins') {
      service = this.pluginService;
    }
    if (name === 'notifications') {
      //service = this.notificationService;
    }
    if (name === 'commands') {
      service = this.commandService;
    }
    if (name === 'datasources') {
      // service = this.datasourceService;
    }
    return service;
  }
}
