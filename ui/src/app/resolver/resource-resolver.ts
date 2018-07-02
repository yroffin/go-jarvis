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
import { ResourceStoreService, SelectResourceAction } from "../store/resources.store";
import { JarvisDefaultResource } from "../interface/jarvis-default-resource";
import { ResourceBean } from "../model/resource-bean";
import { DeviceBean } from "../model/device-bean";

export class ResourceResolver<T extends ResourceBean> implements Resolve<T> {
    private jarvisDataService: JarvisDefaultResource<T>;
    private callback: (resource: any, subject: Subject<any>) => void;

    constructor(
        jarvisDataService: JarvisDefaultResource<T>,
        callback: (resource: any, subject: Subject<any>) => void,
    ) {
        this.jarvisDataService = jarvisDataService;
        this.callback = callback;
    }

    resolve(
        route: ActivatedRouteSnapshot,
        state: RouterStateSnapshot
    ): Observable<any> | Promise<any> | any {
        return this.get(route.params['id'], new BehaviorSubject<any>('select one resource'));
    }

    public get(id: string, subject: Subject<any>): Subject<any> {
        let resource: T;
        this.jarvisDataService.GetSingle(id)
            .subscribe(
            (data: T) => resource = data,
            error => console.log(error),
            () => {
                this.callback(resource, subject);
            }
            );
        return subject;
    }
}