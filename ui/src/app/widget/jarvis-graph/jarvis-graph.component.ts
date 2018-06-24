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

import { Component, OnInit, Input, ViewChild, EventEmitter, Output, ElementRef } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Network, DataSet, Node, Edge, IdType } from 'vis';
import * as _ from 'lodash';
import { AfterViewInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { NodeBean, EdgeBean, GraphBean } from '../../model/graph/graph-bean';
import { OverlayPanel } from 'primeng/overlaypanel';

@Component({
  selector: 'app-jarvis-graph',
  templateUrl: './jarvis-graph.component.html',
  styleUrls: ['./jarvis-graph.component.css']
})
export class JarvisGraphComponent implements OnInit, AfterViewInit {

  protected _network: any;
  protected _graph: GraphBean;
  protected _options: any = {};
  public display = false;

  @ViewChild('wrapper') wrapper: ElementRef;
  @ViewChild('configure') configure: ElementRef;
  @Output() onChange: EventEmitter<any> = new EventEmitter();

  public items: MenuItem[];

  /**
   * internal
   */

  constructor(
  ) {
    this.items = [
      {
        label: 'Configuration',
        command: ($event) => {
          this.display = true;
        }
      }
    ];
  }

  /**
   * init component
   */
  ngOnInit() {
  }

  /**
   * init component
   */
  ngAfterViewInit() {
    setTimeout(() => {
      this.update()
    }, 1000)
  }

  @Input() get options(): any {
    return this._options;
  }

  set options(val: any) {
    if (val) {
      this._options = val;
    } else {
      this._options = {};
    }
  }

  @Input() get graph(): GraphBean {
    return this._graph;
  }

  set graph(val: GraphBean) {
    this._graph = val;
    this.update();
  }

  /**
   * update edge
   */
  public update() {
    if (this._network) {
      this._network.setOptions(this._options);
      if (!this._graph.options) {
        this._graph.options = {};
      }
      this._network.setData(this._graph);
    } else {
      // create a network
      var container = this.wrapper.nativeElement

      this._options.configure = {
        enabled: true,
        container: this.configure.nativeElement
      }

      this._network = new Network(container, this._graph, this._options);

      this._network.on("showPopup", (params) => {
        // Emit event
        this.onChange.emit({ type: "showPopup", params: params });
      });

      this._network.on("selectNode", (params) => {
        // Emit event
        this.onChange.emit({ type: "selectNode", params: params });
      });

      this._network.on("selectEdge", (params) => {
        // Emit event
        this.onChange.emit({ type: "selectEdge", params: params });
      });

      this._network.on("hoverNode", (params) => {
        // Emit event
        this.onChange.emit({ type: "hoverNode", params: params });
      });
    }
  }

}
