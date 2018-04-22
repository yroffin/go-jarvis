/* 
 * Copyright 2017 Yannick Roffin.
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

import { Injectable } from '@angular/core';
import { ActionReducer, Action, State } from '@ngrx/store';
import { Store } from '@ngrx/store';
import { createFeatureSelector, createSelector, MemoizedSelector } from '@ngrx/store';

import * as _ from 'lodash';

import { ActionWithPayload } from './action-with-payload';
import { VersionBean } from '../model/system/version';

/**
 * states
 */
export interface AppState {
  feature: VersionState;
}

export interface VersionState {
  version: VersionBean;
}

/**
 * actions
 */
export class VersionAction implements ActionWithPayload<VersionBean> {
  readonly type = 'VersionAction';
  constructor(public payload: VersionBean) { }
}

export type AllSystemActions = VersionAction;

/**
 * main store for this application
 */
@Injectable()
export class SystemStoreService {

  private getVersion: MemoizedSelector<object, VersionBean>;

  /**
   * 
   * @param _store constructor
   */
  constructor(
    private _store: Store<VersionState>
  ) {
    this.getVersion = createSelector(createFeatureSelector<VersionState>('system'), (state: VersionState) => state.version);
  }

  /**
   * select this store service
   */
  public message(): Store<VersionBean> {
    return this._store.select(this.getVersion);
  }

  /**
   * dispatch
   * @param action dispatch action
   */
  public dispatch(action: AllSystemActions) {
    this._store.dispatch(action);
  }

  /**
   * metareducer (Cf. https://www.concretepage.com/angular-2/ngrx/ngrx-store-4-angular-5-tutorial)
   * @param state 
   * @param action 
   */
  public static reducer(state: VersionState = { version: new VersionBean() }, action: AllSystemActions): VersionState {

    switch (action.type) {
      /**
       * message incomming
       */
      case 'VersionAction':
        {
          let newState = new VersionBean();
          newState.ui = action.payload.ui;
          return {
            version: newState
          };
        }

      default:
        return state;
    }
  }
}
