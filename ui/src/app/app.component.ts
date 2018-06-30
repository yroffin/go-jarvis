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

import { Component, Input, OnInit, ViewChild } from '@angular/core';

import { Observable } from 'rxjs';
import { map, catchError } from 'rxjs/operators';

import { Response, Headers } from '@angular/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Router } from '@angular/router';
import { State, Store } from '@ngrx/store';

import { MatSidenav, MatSnackBar, MatDialog } from '@angular/material';

import { WindowRef } from './service/jarvis-utils.service';
import { JarvisConfigurationService } from './service/jarvis-configuration.service';
import { JarvisSecurityService } from './service/jarvis-security.service';
import { MessageStoreService, Message } from './store/message.store';
import { ProfileGuard } from './guard/profile.service';

/**
 * data model
 */
import { ResourceBean } from './model/resource-bean';
import { Oauth2Bean, MeBean } from './model/security/oauth2-bean';
import { DialogAbout } from './dialog/about/jarvis-about.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  @ViewChild('sidenav') sidenav: MatSidenav;
  myInnerHeight: any;

  /**
   * global system menu
   */
  public items: any[];
  public dispMe: boolean = false;
  public dispHelp: boolean = false;

  protected messageStream: Observable<Message>;

  public me: MeBean;
  public help: string;

  /**
   * constructor
   */
  constructor(
    private profile: ProfileGuard,
    private http: HttpClient,
    private router: Router,
    private windowRef: WindowRef,
    private configuration: JarvisConfigurationService,
    private jarvisSecurityService: JarvisSecurityService,
    private messageStoreService: MessageStoreService,
    private snackBar: MatSnackBar,
    public dialog: MatDialog
  ) {
    this.myInnerHeight = windowRef.getWindow();

    /**
      * register to store
      */
    this.messageStream = this.messageStoreService.message();

    /**
     * register to store update
     */
    this.messageStream
      .subscribe((message) => {
        this.snackBar.open(message.detail, message.severity, {
          duration: 2000,
        });
      });
  }

  /**
   * show me
   */
  showMe() {
    this.me = this.profile.getMe();
    this.dispMe = true;
  }

  /**
   * show me
   */
  showVersion() {
    this.jarvisSecurityService.Version().subscribe(
      (data: any) => {
        const dialogRef = this.dialog.open(DialogAbout, {
          width: '600px',
          data: { payload: JSON.stringify(data, null, '\t') }
        });

        dialogRef.afterClosed().subscribe(result => {
        });
      }
    )
  }

  /**
   * show help
   */
  showHelp() {
    let text;
    let id = 'help' + this.router.routerState.snapshot.url
      .replace(/\//g, "-")
      .replace(/0|1|2|3|4|5|6|7|8|9/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/00/g, "0")
      .replace(/0/g, "by-id");
    this.GetHelp(id + '.markdown').subscribe(
      (data) => {
        text = data;
      },
      (error) => {
        this.help = 'TODO';
      },
      () => {
        this.help = text;
      }
    );
    this.dispHelp = true;
  }

  /**
   * get single resource
   */
  public GetHelp = (id: string): Observable<string> => {
    let headers = new HttpHeaders();
    headers.append('JarvisAuthToken', this.configuration.getJarvisAuthToken());
    return this.http.get(this.configuration.ServerWithApiUrl + 'helps/fr/' + id, { headers: headers }).pipe(
      map((response: Response) => response),
      catchError(this.handleError)
    );
  }

  /**
   * error handler
   */
  protected handleError(error: Response): Observable<any> {
    throw (error || 'Server error');
  }

  /**
   * global init of system menu
   */
  ngOnInit() {
    /**
     * get profile from store
     */
    this.loadMenu();
  }

  /**
   * test if isMobile
   */
  public isMobile(): boolean {
    if (window.matchMedia("(min-width: 400px)").matches) {
      /*
       * the view port is at least 400 pixels wide
       **/
      return false;
    } else {
      /*
       * the view port is less than 400 pixels wide
       **/
      return true;
    }
  }

  /**
   * the given event on the document root.
   */
  private loadMenu(): void {
    /**
     * global item menu
     */
    this.items = [
      {
        label: 'Home',
        icon: 'fa-home',
        routerLink: ['/desktop']
      },
      {
        label: 'Ressource',
        icon: 'fa-sliders',
        items: [
          {
            label: 'Core',
            icon: 'fa-code',
            items: [
              { label: 'View', icon: 'fa-briefcase', routerLink: ['/views'] },
              { label: 'Device', icon: 'fa-server', routerLink: ['/devices'] },
              { label: 'Plugin', icon: 'fa-podcast', routerLink: ['/plugins'] },
              { label: 'Command', icon: 'fa-plug', routerLink: ['/commands'] },
              { label: 'Connector', icon: 'fa-bug', routerLink: ['/connectors'] }
            ]
          },
          {
            label: 'Activation',
            icon: 'fa-sliders',
            items: [
              { label: 'Trigger', icon: 'fa-paper-plane', routerLink: ['/triggers'] },
              { label: 'Cron', icon: 'fa-inbox', routerLink: ['/crons'] },
              { label: 'Notification', icon: 'fa-bolt', routerLink: ['/notifications'] },
              { label: 'Broker', icon: 'fa-exchange', routerLink: ['/broker'] },
            ]
          },
          {
            label: 'Code',
            icon: 'fa-code-fork',
            items: [
              { label: 'Process', icon: 'fa-flash', routerLink: ['/processes'] },
              { label: 'Model', icon: 'fa-briefcase', routerLink: ['/models'] }
            ]
          },
          {
            label: 'Misc',
            icon: 'fa-cube',
            items: [
              { label: 'Configuration', icon: 'fa-database', routerLink: ['/configurations'] },
              { label: 'Snapshot', icon: 'fa-clone', routerLink: ['/snapshots'] },
              { label: 'Property', icon: 'fa-code', routerLink: ['/properties'] },
              { label: 'DataSource', icon: 'fa-podcast', routerLink: ['/datasources'] },
              { label: 'Server', icon: 'fa-line-chart', routerLink: ['/resources'] },
              { label: 'Measure', icon: 'fa-line-chart', routerLink: ['/measures'] }
            ]
          }
        ]
      }
    ];
  }
}
