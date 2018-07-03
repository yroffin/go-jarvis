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

import { Component, OnInit, ViewChild } from '@angular/core';
import { State, Store } from '@ngrx/store';
import * as _ from 'lodash';

import { SelectItem, UIChart } from 'primeng/primeng';
import { JarvisMqttService } from '../../service/jarvis-mqtt.service';

import { Observable } from 'rxjs';
import { MessageBean } from '../../model/broker/message-bean';
import { BrokerStoreService } from '../../store/broker.store';

@Component({
  selector: 'app-jarvis-server-resources',
  templateUrl: './jarvis-server-resources.component.html',
  styleUrls: ['./jarvis-server-resources.component.css']
})
export class JarvisServerResourcesComponent implements OnInit {

  public data: any;
  @ViewChild('chart') chart: UIChart;
  public messageStream: Observable<MessageBean>;

  constructor(
    private mqttService: JarvisMqttService,
    private brokerStoreService: BrokerStoreService
  ) {
    this.data = {
      labels: ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'],
      datasets: []
    }

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
        if (item.topic && !item.topic.startsWith('/system/metrics/')) {
          return;
        }
        let data: any;
        if (item.topic.startsWith('/system/metrics/cpu')) {
          data = item.body[0];
        }
        if (item.topic.startsWith('/system/metrics/mem')) {
          data = item.body;
        }
        _.each(_.keysIn(data), (k) => {
          this.checkDataset(k, data[k]);
        });

        this.chart.refresh();
      });
  }

  private checkDataset(key: string, value: any) {
    let found = false;
    let dataset = _.find(this.data.datasets, (dataset) => {
      return dataset.label === key;
    })
    if (dataset) {
      dataset.data.push(value);
      dataset.data.shift();
    } else {
      let color =  this.getRandomColor();
      this.data.datasets.push({
        label: key,
        data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        fill: false,
        borderColor: color,
        backgroundColor: color
      });
    }
  }

  private getRandomColor() {
    var letters = '0123456789ABCDEF';
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }

}
