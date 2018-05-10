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

import { Component, OnInit, Input, ViewChild, EventEmitter, Output } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Network, DataSet, Node, Edge, IdType } from 'vis';
import * as _ from 'lodash';
import { AfterViewInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { NodeBean, EdgeBean } from '../../model/graph/graph-bean';
import { OverlayPanel } from 'primeng/overlaypanel';

@Component({
  selector: 'app-jarvis-graph',
  templateUrl: './jarvis-graph.component.html',
  styleUrls: ['./jarvis-graph.component.css']
})
export class JarvisGraphComponent implements OnInit, AfterViewInit {

  protected _nodes: NodeBean[];
  protected _edges: EdgeBean[];
  protected _options: any;
  public display = false;

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
    this._options = val;
  }

  @Input() get nodes(): NodeBean[] {
    return this._nodes;
  }

  set nodes(val: NodeBean[]) {
    this._nodes = val;
  }

  @Input() get edges(): EdgeBean[] {
    return this._edges;
  }

  set edges(val: EdgeBean[]) {
    this._edges = val;
  }

  /**
   * update edge
   */
  public update() {
    // create a network
    var container = document.getElementById('mynetwork');
    var data = {
      nodes: this._nodes,
      edges: this._edges
    };

    this._options.configure = {
      enabled: true,
      container: document.getElementById('myconfigure')
    }

    var network = new Network(container, data, this._options);

    network.on("showPopup", (params) => {
      // Emit event
      this.onChange.emit({ type: "showPopup", params: params });
    });

    network.on("selectNode", (params) => {
      // Emit event
      this.onChange.emit({ type: "selectNode", params: params });
    });

    network.on("selectEdge", (params) => {
      // Emit event
      this.onChange.emit({ type: "selectEdge", params: params });
    });

    network.on("hoverNode", (params) => {
      // Emit event
      this.onChange.emit({ type: "hoverNode", params: params });
    });
  }

}
