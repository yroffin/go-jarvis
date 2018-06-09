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

import { Component, AfterContentInit, Input, ViewChild, EventEmitter, Output, ElementRef, ChangeDetectorRef, NgZone } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Network, DataSet, Node, Edge, IdType } from 'vis';
import * as _ from 'lodash';
import { AfterViewInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { NodeBean, EdgeBean } from '../../model/graph/graph-bean';
import { OverlayPanel } from 'primeng/overlaypanel';

declare var BpmnJS: any;

@Component({
  selector: 'app-jarvis-bpmnjs',
  templateUrl: './jarvis-bpmnjs.component.html',
  styleUrls: ['./jarvis-bpmnjs.component.css']
})
export class JarvisBpmnComponent implements AfterContentInit {

  @ViewChild('wrapper') wrapper: ElementRef;
  protected _bpmn: string;
  protected _display: boolean = false;
  private viewer: any;

  @Output() onChange: EventEmitter<any> = new EventEmitter();

  /**
   * internal
   */

  constructor(
    private cd: ChangeDetectorRef,
    private zone: NgZone
  ) {
  }

  /**
   * after init component
   */
  ngAfterContentInit() {
    // BPMN
    this.viewer = new BpmnJS({
      container: this.wrapper.nativeElement
    });
  }

  init() {
    this.update();
  }

  @Input() get display(): any {
    return this._display;
  }

  set display(val: any) {
    this._display = val;
  }

  @Input() get bpmn(): any {
    return this._bpmn;
  }

  set bpmn(val: any) {
    this._bpmn = val;
    this.update();
  }

  /**
   * update edge
   */
  public update() {
    if (this.viewer) {
      this.viewer.importXML(this._bpmn, (err) => {
        if (!err) {
          this.viewer.get('canvas').zoom('fit-viewport');
        } else {
          console.log('something went wrong:', err);
        }
      });
    }
  }

}
