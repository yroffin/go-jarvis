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

import { Router, ActivatedRoute, Params } from '@angular/router';

import { Observable } from 'rxjs';
import { map, catchError } from 'rxjs/operators';

import { HttpClient, HttpHeaders } from '@angular/common/http';

import { JarvisDefaultResource } from '../interface/jarvis-default-resource';
import { JarvisSecurityService } from '../service/jarvis-security.service';
import { JarvisConfigurationService } from '../service/jarvis-configuration.service';

/**
 * data model
 */
import { ResourceBean } from '../model/resource-bean';

/**
 * default class to handle default behaviour or resource
 * component
 */
export class JarvisDataCoreResource<T extends ResourceBean> implements JarvisDefaultResource<T> {

    protected actionUrl: string;
    protected headers: HttpHeaders;
    protected http: HttpClient;
    protected configuration: JarvisConfigurationService;

    /**
     * constructor
     */
    constructor(_configuration: JarvisConfigurationService, actionUrl: string, _http: HttpClient) {
        this.http = _http;
        this.actionUrl = actionUrl;
        this.configuration = _configuration;

        this.headers = new HttpHeaders();
        this.headers.append('Content-Type', 'application/json');
        this.headers.append('Accept', 'application/json');
        this.headers.append('JarvisAuthToken', this.configuration.getJarvisAuthToken());
    }

    /**
     * execute remote task on this resource
     */
    public Task<R>(id: string, task: string, args: any): Observable<R> {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.post(this.actionUrl + '/' + id + '?task=' + task, JSON.stringify(args), { headers: this.headers }).pipe(
            map((response: Response) => {
                return response
            }),
            catchError(this.handleError));
    }

    /**
     * execute remote task on this resource
     */
    public TaskAsXml = (id: string, task: string, args: any): Observable<any> => {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.post(this.actionUrl + '/' + id + '?task=' + task, JSON.stringify(args), { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * get all resources
     */
    public GetAll = (): Observable<T[]> => {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.get(this.actionUrl, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * get single resource
     */
    public GetSingle = (id: string): Observable<T> => {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.get(this.actionUrl + '/' + id, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * add a new resource
     */
    public Add = (itemToAdd: T): Observable<T> => {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.post(this.actionUrl, JSON.stringify(itemToAdd), { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * update this resource
     */
    public Update = (id: string, itemToUpdate: T): Observable<T> => {
        this.headers.append('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.put(this.actionUrl + '/' + id, JSON.stringify(itemToUpdate), { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * delete this resource
     */
    public Delete = (id: string): Observable<T> => {
        this.headers.set('JarvisAuthToken', this.configuration.getJarvisAuthToken());
        return this.http.delete(this.actionUrl + '/' + id, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * error handler
     */
    protected handleError(error: Response): Observable<any> {
        throw(error || 'Server error');
    }
}
