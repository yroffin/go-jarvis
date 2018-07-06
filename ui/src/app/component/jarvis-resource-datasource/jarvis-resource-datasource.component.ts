/* 
 * Copyright 2017 Yannick Roffin.
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

import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import * as _ from 'lodash';

import { Router, ActivatedRoute, Params } from '@angular/router';
import { SelectItem, UIChart } from 'primeng/primeng';

import { JarvisConfigurationService } from '../../service/jarvis-configuration.service';
import { JarvisResourceLink } from '../../class/jarvis-resource-link';

import { JarvisDataDatasourceService } from '../../service/jarvis-data-datasource.service';
import { LoggerService } from '../../service/logger.service';

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
import { DataSourceBean } from '../../model/connector/datasource-bean';
import { MatTableDataSource } from '@angular/material';
import { JarvisPickResourceService } from '../../service/jarvis-pick-resource.service';
import { ResourceStoreService } from '../../store/resources.store';

@Component({
  selector: 'app-jarvis-resource-datasource',
  templateUrl: './jarvis-resource-datasource.component.html',
  styleUrls: ['./jarvis-resource-datasource.component.css']
})
export class JarvisResourceDatasourceComponent extends JarvisResource<DataSourceBean> implements NotifyCallback<ResourceBean>, OnInit {

  /**
   * internal
   */
  myStream: Observable<DataSourceBean>;
  @Input() myDataSource: DataSourceBean;
  public myValues: string[];

  /**
   * constructor
   */
  constructor(
    private _route: ActivatedRoute,
    private _router: Router,
    private jarvisDataDatasourceService: JarvisDataDatasourceService,
    private resourceStoreService: ResourceStoreService,
    private jarvisPickResourceService: JarvisPickResourceService) {
    super('/datasources', [], jarvisDataDatasourceService, _route, _router);
    this.myStream = this.resourceStoreService.datasource();
  }

  /**
   * load device and related data
   */
  ngOnInit() {
    this.myStream.subscribe(
      (resource: DataSourceBean) => {
        this.setResource(this.myDataSource);
        this.myDataSource = <DataSourceBean>resource;
        this.values();
      }
    )
  }

  /**
   * task action
   */
  public values(): void {
    this.jarvisDataDatasourceService.Task('*', 'values', {})
      .subscribe(
      (result: any) => this.myValues = result,
      error => console.log(error),
      () => {
      });
  }

  /**
   * notify to add new resource
   */
  public notify(picker: PickerBean, resource: ResourceBean): void {
  }

  /**
   * pick datasources
   */
  public pick(picker: PickerBean): void {
  }
}
