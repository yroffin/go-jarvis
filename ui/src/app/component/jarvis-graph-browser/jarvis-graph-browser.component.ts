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

import { Component, Input, OnInit } from '@angular/core';
import * as _ from 'lodash';
import { NodeBean, EdgeBean, GraphBean } from '../../model/graph/graph-bean';
import { Store } from '@ngrx/store/src/store';
import { GraphStoreService, LoadGraphAction } from '../../store/graph.store';
import { JarvisDataSnapshotService } from '../../service/jarvis-data-snapshot.service';

@Component({
  selector: 'app-jarvis-graph-browser',
  templateUrl: './jarvis-graph-browser.component.html',
  styleUrls: ['./jarvis-graph-browser.component.css']
})
export class JarvisGraphBrowserComponent implements OnInit {

  public options: any;

  protected graphStream: Store<GraphBean>;
  public graph: GraphBean;

  constructor(
    private graphStoreService: GraphStoreService,
    private jarvisDataSnapshotService: JarvisDataSnapshotService
  ) {
    this.graphStream = this.graphStoreService.graph()
  }

  ngOnInit() {
    this.loadGraph()

    this.graphStream.subscribe(
      (element: GraphBean) => {
        this.graph = element;
        this.options = element.options;
      },
      error => {
        console.error(error);
      },
      () => {
      });
  }

  loadGraph() {
    this.jarvisDataSnapshotService.Task<GraphBean>('*', 'graph', {})
      .subscribe(
      (data) => {
        // fix graph data
        this.graphStoreService.dispatch(new LoadGraphAction(data))
      },
      (error: any) => {
        console.error("Error while loading graph");
      },
      () => {
      });
  }

  handler(event: any) {
    console.warn("event", event);
  }
}
