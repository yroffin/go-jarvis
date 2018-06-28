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

import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import * as _ from 'lodash';
import { State, Store } from '@ngrx/store';

import { SelectItem, UIChart } from 'primeng/primeng';

import { Observable } from 'rxjs';
import { MessageBean } from '../../model/broker/message-bean';
import { BrokerStoreService } from '../../store/broker.store';
import { MatTableDataSource, MatPaginator } from '@angular/material';
import { JarvisMqttService } from '../../service/jarvis-mqtt.service';
import { AfterViewInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { LoggerService } from '../../service/logger.service';

@Component({
  selector: 'app-jarvis-broker',
  templateUrl: './jarvis-broker.component.html',
  styleUrls: ['./jarvis-broker.component.css']
})
export class JarvisBrokerComponent implements OnInit, AfterViewInit {

  @ViewChild('paginator') paginator: MatPaginator;
  @Input() myMessages: MessageBean[] = <MessageBean[]>[];
  public myMatResources = new MatTableDataSource([])
  public messageStream: Observable<MessageBean>;

  private ids: number = 0;

  constructor(
    private store: Store<State<MessageBean>>,
    private mqttService: JarvisMqttService,
    private brokerStoreService: BrokerStoreService,
    private logger: LoggerService
  ) {
    /**
      * register to store
      */
    this.messageStream = this.brokerStoreService.message();
  }

  ngOnInit() {
    /**
     * register to store update
     */
    this.messageStream
      .subscribe((item) => {
        let msg = new MessageBean();
        msg.id = '' + (this.ids++);
        msg.topic = item.topic;
        try {
          msg.body = JSON.stringify(item.body, null, '\t');
        } catch (Exc) {
          this.logger.warn("Bad body (not json");
          msg.body = item.body;
        }
        if (this.myMessages.length > 256) {
          this.myMessages.shift()
        }
        this.myMessages.push(msg);
        this.myMatResources.data = this.myMessages
      });
  }

  /**
 * Set the paginator after the view init since this component will
 * be able to query its view for the initialized paginator.
 */
  ngAfterViewInit() {
    this.myMatResources.paginator = this.paginator;
  }

}
