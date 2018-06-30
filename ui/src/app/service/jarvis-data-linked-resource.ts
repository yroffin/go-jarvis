import { Router, ActivatedRoute, Params } from '@angular/router';

import { Observable } from 'rxjs';
import { map, catchError } from 'rxjs/operators';

import { Response } from '@angular/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { JarvisDefaultResource, JarvisDefaultLinkResource } from '../interface/jarvis-default-resource';

/**
 * data model
 */
import { ResourceBean } from '../model/resource-bean';

/**
 * default class to handle default behaviour or resource
 * component
 */
export class JarvisDataLinkedResource<T extends ResourceBean> implements JarvisDefaultLinkResource<T> {

    private actionUrl: string;
    private headers: HttpHeaders;
    private link: string;
    private http: HttpClient;

    /**
     * constructor
     */
    constructor(actionUrl: string, link: string, _http: HttpClient) {
        this.http = _http;
        this.link = link;
        this.actionUrl = actionUrl;

        this.headers = new HttpHeaders();
        this.headers.append('Content-Type', 'application/json');
        this.headers.append('Accept', 'application/json');
    }

    /**
     * get all link
     */
    public GetAll = (id: string): Observable<T[]> => {
        return this.http.get(this.actionUrl + '/' + id + this.link, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * get all link
     */
    public FindAll = (id: string, filters: string): Observable<T[]> => {
        return this.http.get(this.actionUrl + '/' + id + this.link + '?' + filters, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * get single link
     */
    public GetSingle = (id: string, linkId: string): Observable<T> => {
        return this.http.get(this.actionUrl + '/' + id  + this.link + '/' + linkId, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * add a new link
     */
    public Add = (id: string, linkId: string, linkToAdd: any): Observable<T> => {
        return this.http.post(this.actionUrl + '/' + id  + this.link + '/' + linkId, JSON.stringify(linkToAdd), { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * update a link
     */
    public Update = (id: string, linkId: string, instance: string, linkToUpdate: any): Observable<T> => {
        return this.http.put(this.actionUrl + '/' + id  + this.link + '/' + linkId + '?instance=' + instance, JSON.stringify(linkToUpdate), { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * delete a link
     */
    public Delete = (id: string, linkId: string, instance: string): Observable<T> => {
        return this.http.delete(this.actionUrl + '/' + id  + this.link + '/' + linkId + '?instance=' + instance, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * delete a link
     */
    public DeleteWithFilter = (id: string, linkId: string, instance: string, filters: string): Observable<T> => {
        return this.http.delete(this.actionUrl + '/' + id  + this.link + '/' + linkId + '?instance=' + instance + '&' + filters, { headers: this.headers }).pipe(
            map((response: Response) => response),
            catchError(this.handleError));
    }

    /**
     * handle error
     */
    private handleError(error: Response): Observable<any> {
        throw(error || 'Server error');
    }
}
