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
import { NodeBean, EdgeBean, GraphBean } from '../model/graph/graph-bean';

/**
 * states
 */
export interface AppState {
    feature: GraphState;
}

export interface GraphState {
    graph: GraphBean;
}

/**
 * actions
 */
export class LoadGraphAction implements ActionWithPayload<GraphBean> {
    readonly type = 'LoadGraphAction';
    constructor(public payload: GraphBean) { }
}

export type AllGraphActions = LoadGraphAction;

/**
 * main store for this application
 */
@Injectable()
export class GraphStoreService {

    private getGraph: MemoizedSelector<object, GraphBean>;

    /**
     * 
     * @param _store constructor
     */
    constructor(
        private _store: Store<GraphState>
    ) {
        this.getGraph = createSelector(createFeatureSelector<GraphState>('graph'), (state: GraphState) => state.graph);
    }

    /**
     * select this store service
     */
    public graph(): Store<GraphBean> {
        return this._store.select(this.getGraph);
    }

    /**
     * dispatch
     * @param action dispatch action
     */
    public dispatch(action: AllGraphActions) {
        this._store.dispatch(action);
    }

    /**
     * metareducer (Cf. https://www.concretepage.com/angular-2/ngrx/ngrx-store-4-angular-5-tutorial)
     * @param state 
     * @param action 
     */
    public static reducer(state: GraphState = { graph: new GraphBean() }, action: AllGraphActions): GraphState {

        switch (action.type) {
            /**
             * message incomming
             */
            case 'LoadGraphAction':
                {
                    let graph = action.payload;

                    return {
                        graph: action.payload
                    };
                }

            default:
                return state;
        }
    }
}