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

import { Component, Input, ViewChild, OnInit, AfterContentInit, ElementRef } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Observable } from 'rxjs';

declare var Prism: any;

import { LoggerService } from '../../service/logger.service';
import { JarvisDataTriggerService } from '../../service/jarvis-data-trigger.service';
import { JarvisDataProcessService } from '../../service/jarvis-data-process.service';
import { JarvisResourceLink } from '../../class/jarvis-resource-link';

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
import { ProcessBean } from '../../model/code/process-bean';
import { TriggerBean } from '../../model/trigger-bean';
import { JarvisBpmnComponent } from '../../widget/jarvis-bpmnjs/jarvis-bpmnjs.component';
import { GraphBean, NodeBean, EdgeBean } from '../../model/graph/graph-bean';
import { LoadGraphAction, GraphStoreService } from '../../store/graph.store';
import { Store } from '@ngrx/store/src/store';
import { JarvisGraphExplorerComponent } from '../../widget/jarvis-graph-explorer/jarvis-graph-explorer.component';
import { MatTableDataSource } from '@angular/material';
import { JarvisPickResourceService } from '../../service/jarvis-pick-resource.service';
import { ResourceStoreService } from '../../store/resources.store';

@Component({
  selector: 'app-jarvis-resource-process',
  templateUrl: './jarvis-resource-process.component.html',
  styleUrls: ['./jarvis-resource-process.component.css']
})
export class JarvisResourceProcessComponent extends JarvisResource<ProcessBean> implements NotifyCallback<ResourceBean>, OnInit, AfterContentInit {

  myStream: Observable<ProcessBean>;
  @Input() myProcess: ProcessBean;
  public myMatResources = new MatTableDataSource([]);

  @ViewChild('wrapBpmn') wrapBpmn: JarvisBpmnComponent;
  @ViewChild('wrapGraph') wrapGraph: JarvisGraphExplorerComponent;

  /**
   * internal data
   */
  private display: boolean = false;
  private myData: any = {};
  private myOutputData: any = {};
  private myTrigger: TriggerBean;
  private jarvisTriggerLink: JarvisResourceLink<TriggerBean>;

  protected graphStream: Observable<GraphBean>;
  public graph: GraphBean;

  /**
   * constructor
   */
  constructor(
    private _route: ActivatedRoute,
    private _router: Router,
    private logger: LoggerService,
    private processService: JarvisDataProcessService,
    private _triggerService: JarvisDataTriggerService,
    private graphStoreService: GraphStoreService,
    private resourceStoreService: ResourceStoreService,
    private jarvisPickResourceService: JarvisPickResourceService,
  ) {
    super('/processes', ['deploy', 'test'], processService, _route, _router);
    this.jarvisTriggerLink = new JarvisResourceLink<TriggerBean>(this.logger);
    this.graphStream = this.graphStoreService.graph()
    this.myStream = this.resourceStoreService.process();
  }

  ngOnInit() {
  }

  ngAfterContentInit() {
    this.myStream.subscribe(
      (resource: ProcessBean) => {
        this.setResource(resource);
        let picker: PickerBean = new PickerBean();
        picker.action = 'complete';
        this.notify(picker, resource);
      }
    )

    this.graphStream.subscribe(
      (graph: GraphBean) => {
        if(this.wrapGraph) {
          this.wrapGraph.graph = graph;
        }
      },
      error => {
        console.error(error);
      },
      () => {
      });
  }

  /**
   * highlight source
   */
  public hightlight(body: string): void {
    if (body) {
      return Prism.highlight(body, Prism.languages.xml);
    }
  }

  /**
   * pick bpmn
   */
  bpmn() {
    this.wrapBpmn.bpmn = this.myProcess.bpmn;
    this.wrapBpmn.display = true;
  }

  /**
   * pick action
   */
  public pick(picker: PickerBean): void {
    /**
     * find notifications
     */
    if (picker.action === 'triggers') {
      this.jarvisPickResourceService.open(this, 'triggers', picker);
    }
  }

  /**
   * task action
   */
  public deploy(): void {
    this.processService.Update(this.myProcess.id, this.myProcess)
      .subscribe(
      (data: ProcessBean) => data,
      error => console.log(error),
      () => {
        this.logger.info("Save", this.myProcess);
        /**
         * execute this plugin
         */
        let myOutputData;
        this.processService.Task(this.myProcess.id, 'deploy', {})
          .subscribe(
          (result: any) => myOutputData = result,
          error => console.log(error),
          () => {
            this.logger.info("Deploy", this.myProcess);
          }
          );
      });
  }

  /**
   * task action
   */
  public test(): void {
    /**
     * execute this plugin
     */
    let myOutputData;
    this.processService.Task(this.myProcess.id, 'execute', {})
      .subscribe(
      (result: any) => myOutputData = result,
      error => console.log(error),
      () => {
      }
      );
  }

  /**
   * graph action
   */
  public refresh(): void {
    this.processService.Task<GraphBean>('*', 'active', {})
      .subscribe(
      (data) => {
        // fix graph data
        this.graphStoreService.dispatch(new LoadGraphAction(data))
        this.wrapGraph.display = true;
      },
      (error: any) => {
        console.error("Error while loading graph");
      },
      () => {
      });
  }

  /**
   * notify to add new resource
   */
  public notify(picker: PickerBean, resource: ResourceBean): void {
    if (picker.action === 'triggers') {
      this.jarvisTriggerLink.addLink(this.getResource().id, resource.id, this.getResource().triggers, { "order": "1", href: "HREF" }, this.processService.allLinkedTrigger);
    }
    if (picker.action === 'complete') {
      this.myProcess = <ProcessBean>resource;
      this.myProcess.triggers = [];
      (new JarvisResourceLink<TriggerBean>(this.logger)).loadLinksWithCallback(resource.id, this.myProcess.triggers, this.processService.allLinkedTrigger, (elements) => {
        this.myProcess.triggers = elements;
        this.myMatResources = new MatTableDataSource(elements);
      });
    }
  }

  /**
   * drop link
   */
  public dropTriggerLink(linked: TriggerBean): void {
    this.jarvisTriggerLink.dropLink(linked, this.myProcess.id, this.myProcess.triggers, this.processService.allLinkedTrigger);
  }

  /**
   * drop link
   */
  public updateTriggerLink(linked: TriggerBean): void {
    this.jarvisTriggerLink.updateLink(linked, this.myProcess.id, this.processService.allLinkedTrigger);
  }

  /**
   * goto link
   */
  public gotoTriggerLink(linked: TriggerBean): void {
    this._router.navigate(['/triggers/' + linked.id]);
  }
}
