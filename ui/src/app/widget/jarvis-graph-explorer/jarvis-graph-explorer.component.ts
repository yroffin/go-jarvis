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

import { Component, AfterContentInit, Input, ViewChild, EventEmitter, Output, ElementRef, AfterViewInit, ChangeDetectorRef, NgZone } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Network, DataSet, Node, Edge, IdType } from 'vis';
import * as _ from 'lodash';
import { NodeBean, EdgeBean, GraphBean } from '../../model/graph/graph-bean';

@Component({
  selector: 'app-jarvis-graph-explorer',
  templateUrl: './jarvis-graph-explorer.component.html',
  styleUrls: ['./jarvis-graph-explorer.component.css']
})
export class JarvisGraphExplorerComponent implements AfterContentInit {

  @ViewChild('wrapper') wrapper: ElementRef;
  protected _graph: GraphBean;
  public _display: boolean = false;

  public nodes: NodeBean[] = [];
  public edges: EdgeBean[] = [];
  public options: any = {};


  @Output() onChange: EventEmitter<any> = new EventEmitter();

  /**
   * internal
   */

  constructor(
  ) {
  }

  /**
   * after init component
   */
  ngAfterContentInit() {
  }

  init() {
  }

  @Input() get display(): any {
    return this._display;
  }

  set display(val: any) {
    this._display = val;
  }

  @Input() get graph(): any {
    return this._graph;
  }

  set graph(val: any) {
    if (val) {
      this._graph = val;
      this.nodes = this._graph.nodes;
      this.edges = this._graph.edges;
      this.options = this._graph.options;
      console.error("graph",this.nodes)
    }
  }

  handler(event: any) {
    console.warn("event", event);
  }
}
