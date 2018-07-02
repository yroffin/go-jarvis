/* 
 * Copyright 2018 Yannick Roffin.
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
import { Injectable } from "@angular/core";
import { Resolve, RouterStateSnapshot, ActivatedRouteSnapshot } from "@angular/router";
import { Observable, BehaviorSubject, Subject } from "rxjs";
import { ResourceStoreService, GetSnapshotAction, SelectResourceAction } from "../store/resources.store";
import { ResourceResolver } from "./resource-resolver";
import { SnapshotBean } from "../model/misc/snapshot-bean";
import { JarvisDataSnapshotService } from "../service/jarvis-data-snapshot.service";

@Injectable({
    providedIn: 'root'
})
export class SnapshotResolver extends ResourceResolver<SnapshotBean> {
    constructor(
        private jarvisDataSnapshotService: JarvisDataSnapshotService,
        private resourceStoreService: ResourceStoreService,
    ) {
        super(jarvisDataSnapshotService, (resource: any, subject: Subject<any>) => {
            this.resourceStoreService.dispatch(new GetSnapshotAction(resource, subject));
        });
    }
}