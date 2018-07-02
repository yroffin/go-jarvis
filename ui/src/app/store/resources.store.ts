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
import { Observable, Subject } from 'rxjs';

import { ActionReducer, Action, State } from '@ngrx/store';
import { Store } from '@ngrx/store';
import { createFeatureSelector, createSelector, Selector } from '@ngrx/store';

import * as _ from 'lodash';

import { ActionWithPayload, ActionWithPayloadAndSubject } from './action-with-payload';
import { DeviceBean } from '../model/device-bean';
import { ResourceBean } from '../model/resource-bean';
import { MeasureBean } from '../model/connector/measure-bean';
import { CommandBean } from '../model/command-bean';
import { ConfigurationBean } from '../model/configuration-bean';
import { ConnectorBean } from '../model/connector/connector-bean';
import { CronBean } from '../model/cron-bean';
import { DataSourceBean } from '../model/connector/datasource-bean';
import { ModelBean } from '../model/code/model-bean';
import { NotificationBean } from '../model/notification-bean';
import { PluginBean } from '../model/plugin-bean';
import { ProcessBean } from '../model/code/process-bean';
import { SnapshotBean } from '../model/misc/snapshot-bean';
import { TriggerBean } from '../model/trigger-bean';
import { ViewBean } from '../model/view-bean';
import { PropertyBean } from '../model/property-bean';

/**
 * states
 */
export interface AppState {
    feature: ResourceState;
}

export interface ResourceState {
    devices: Array<DeviceBean>;
    device: DeviceBean;
    measure: MeasureBean;
    command: CommandBean;
    configuration: ConfigurationBean;
    connector: ConnectorBean;
    cron: CronBean;
    datasource: DataSourceBean;
    model: ModelBean;
    notification: NotificationBean;
    plugin: PluginBean;
    process: ProcessBean;
    snapshot: SnapshotBean;
    trigger: TriggerBean;
    view: ViewBean;
    property: PropertyBean;
}

/**
 * actions
 */
export class LoadDevicesAction implements ActionWithPayloadAndSubject<Array<DeviceBean>> {
    readonly type = 'LoadDevicesAction';
    public subject: Subject<any>;
    constructor(public payload: Array<DeviceBean>, subject: Subject<any>) { }
}

export class SelectResourceAction<T extends ResourceBean> extends ActionWithPayloadAndSubject<T> {
    constructor(public payload: T, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetDeviceAction extends SelectResourceAction<DeviceBean> {
    readonly type = 'GetDeviceAction';
    constructor(public payload: DeviceBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetMeasureAction extends SelectResourceAction<MeasureBean> {
    readonly type = 'GetMeasureAction';
    constructor(public payload: MeasureBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetCommandAction extends SelectResourceAction<CommandBean> {
    readonly type = 'GetCommandAction';
    constructor(public payload: CommandBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetConfigurationAction extends SelectResourceAction<ConfigurationBean> {
    readonly type = 'GetConfigurationAction';
    constructor(public payload: ConfigurationBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetConnectorAction extends SelectResourceAction<ConnectorBean> {
    readonly type = 'GetConnectorAction';
    constructor(public payload: ConnectorBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetCronAction extends SelectResourceAction<CronBean> {
    readonly type = 'GetCronAction';
    constructor(public payload: CronBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetDatasourceAction extends SelectResourceAction<DataSourceBean> {
    readonly type = 'GetDatasourceAction';
    constructor(public payload: DataSourceBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetModelAction extends SelectResourceAction<ModelBean> {
    readonly type = 'GetModelAction';
    constructor(public payload: ModelBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetNotificationAction extends SelectResourceAction<NotificationBean> {
    readonly type = 'GetNotificationAction';
    constructor(public payload: NotificationBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetPluginAction extends SelectResourceAction<PluginBean> {
    readonly type = 'GetPluginAction';
    constructor(public payload: PluginBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetProcessAction extends SelectResourceAction<ProcessBean> {
    readonly type = 'GetProcessAction';
    constructor(public payload: ProcessBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetSnapshotAction extends SelectResourceAction<SnapshotBean> {
    readonly type = 'GetSnapshotAction';
    constructor(public payload: SnapshotBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetTriggerAction extends SelectResourceAction<TriggerBean> {
    readonly type = 'GetTriggerAction';
    constructor(public payload: TriggerBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetViewAction extends SelectResourceAction<ViewBean> {
    readonly type = 'GetViewAction';
    constructor(public payload: ViewBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export class GetPropertyAction extends SelectResourceAction<PropertyBean> {
    readonly type = 'GetPropertyAction';
    constructor(public payload: ViewBean, subject: Subject<any>) {
        super(payload, subject);
    }
}

export type AllResourcesActions = LoadDevicesAction
    | GetDeviceAction
    | GetMeasureAction
    | GetCommandAction
    | GetConfigurationAction
    | GetConnectorAction
    | GetCronAction
    | GetDatasourceAction
    | GetModelAction
    | GetNotificationAction
    | GetPluginAction
    | GetProcessAction
    | GetSnapshotAction
    | GetTriggerAction
    | GetPropertyAction
    | GetViewAction;

/**
 * main store for this application
 */
@Injectable({
    providedIn: 'root'
})
export class ResourceStoreService {

    private getDevices: Selector<object, Array<DeviceBean>>;
    private getDevice: Selector<object, DeviceBean>;
    private getMeasure: Selector<object, MeasureBean>;
    private getView: Selector<object, ViewBean>;
    private getPlugin: Selector<object, PluginBean>;
    private getCommand: Selector<object, CommandBean>;
    private getConnector: Selector<object, ConnectorBean>;
    private getTrigger: Selector<object, TriggerBean>;
    private getCron: Selector<object, CronBean>;
    private getNotification: Selector<object, NotificationBean>;
    private getProcess: Selector<object, ProcessBean>;
    private getModel: Selector<object, ModelBean>;
    private getConfiguration: Selector<object, ConfigurationBean>;
    private getDatasource: Selector<object, DataSourceBean>;
    private getSnapshot: Selector<object, SnapshotBean>;
    private getProperty: Selector<object, PropertyBean>;

    /**
     * 
     * @param _store constructor
     */
    constructor(
        private _store: Store<ResourceState>
    ) {
        this.getDevices = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.devices);
        this.getDevice = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.device);
        this.getMeasure = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.measure);
        this.getView = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.view);
        this.getPlugin = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.plugin);
        this.getCommand = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.command);
        this.getConnector = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.connector);
        this.getTrigger = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.trigger);
        this.getCron = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.cron);
        this.getNotification = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.notification);
        this.getProcess = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.process);
        this.getModel = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.model);
        this.getConfiguration = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.configuration);
        this.getDatasource = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.datasource);
        this.getSnapshot = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.snapshot);
        this.getProperty = createSelector(createFeatureSelector<ResourceState>('resources'), (state: ResourceState) => state.property);
    }

    /**
     * select this store service
     */
    public devices(): Observable<Array<DeviceBean>> {
        return this._store.select(this.getDevices);
    }

    /**
     * select this store service
     */
    public device(): Observable<DeviceBean> {
        return this._store.select(this.getDevice);
    }

    /**
     * select this store service
     */
    public measure(): Observable<MeasureBean> {
        return this._store.select(this.getMeasure);
    }

    /**
     * select this store service
     */
    public view(): Observable<ViewBean> {
        return this._store.select(this.getView);
    }

    /**
     * select this store service
     */
    public plugin(): Observable<PluginBean> {
        return this._store.select(this.getPlugin);
    }

    /**
     * select this store service
     */
    public command(): Observable<CommandBean> {
        return this._store.select(this.getCommand);
    }

    /**
     * select this store service
     */
    public connector(): Observable<ConnectorBean> {
        return this._store.select(this.getConnector);
    }

    /**
     * select this store service
     */
    public trigger(): Observable<TriggerBean> {
        return this._store.select(this.getTrigger);
    }

    /**
     * select this store service
     */
    public cron(): Observable<CronBean> {
        return this._store.select(this.getCron);
    }

    /**
     * select this store service
     */
    public notification(): Observable<NotificationBean> {
        return this._store.select(this.getNotification);
    }

    /**
     * select this store service
     */
    public process(): Observable<ProcessBean> {
        return this._store.select(this.getProcess);
    }

    /**
     * select this store service
     */
    public model(): Observable<ModelBean> {
        return this._store.select(this.getModel);
    }

    /**
     * select this store service
     */
    public configuration(): Observable<ConfigurationBean> {
        return this._store.select(this.getConfiguration);
    }

    /**
     * select this store service
     */
    public snapshot(): Observable<SnapshotBean> {
        return this._store.select(this.getSnapshot);
    }

    /**
     * select this store service
     */
    public property(): Observable<PropertyBean> {
        return this._store.select(this.getProperty);
    }

    /**
     * select this store service
     */
    public datasource(): Observable<DataSourceBean> {
        return this._store.select(this.getDatasource);
    }

    /**
     * dispatch
     * @param action dispatch action
     */
    public dispatch(action: AllResourcesActions) {
        this._store.dispatch(action);
    }

    /**
     * metareducer (Cf. https://www.concretepage.com/angular-2/ngrx/ngrx-store-4-angular-5-tutorial)
     * @param state 
     * @param action 
     */
    public static reducer(state: ResourceState = {
        devices: new Array<DeviceBean>(),
        device: new DeviceBean(),
        measure: new MeasureBean(),
        command: new CommandBean(),
        configuration: new ConfigurationBean(),
        connector: new ConnectorBean(),
        cron: new CronBean(),
        datasource: new DataSourceBean(),
        model: new ModelBean(),
        notification: new NotificationBean(),
        plugin: new PluginBean(),
        process: new ProcessBean(),
        snapshot: new SnapshotBean(),
        trigger: new TriggerBean(),
        view: new ViewBean(),
        property: new PropertyBean(),
    }, action: AllResourcesActions): ResourceState {

        switch (action.type) {
            /**
             * message incomming
             */
            case 'LoadDevicesAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.devices = action.payload;
                    });
                }

            case 'GetDeviceAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.device = action.payload;
                    });
                }

            case 'GetViewAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.view = action.payload;
                    });
                }

            case 'GetPluginAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.plugin = action.payload;
                    });
                }

            case 'GetCommandAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.command = action.payload;
                    });
                }

            case 'GetConnectorAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.connector = action.payload;
                    });
                }

            case 'GetTriggerAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.trigger = action.payload;
                    });
                }

            case 'GetCronAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.cron = action.payload;
                    });
                }

            case 'GetNotificationAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.notification = action.payload;
                    });
                }

            case 'GetProcessAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.process = action.payload;
                    });
                }

            case 'GetModelAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.model = action.payload;
                    });
                }

            case 'GetConfigurationAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.configuration = action.payload;
                    });
                }

            case 'GetSnapshotAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.snapshot = action.payload;
                    });
                }

            case 'GetPropertyAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.property = action.payload;
                    });
                }

            case 'GetDatasourceAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.datasource = action.payload;
                    });
                }

            case 'GetMeasureAction':
                {
                    return ResourceStoreService.resolve(action, state, (state) => {
                        state.measure = action.payload;
                    });
                }

            default:
                return state;
        }
    }

    private static resolve(action: AllResourcesActions, state: ResourceState, callback: (ResourceState) => void): any {
        action.subject.complete();
        let resolved = {
            devices: state.devices,
            device: state.device,
            measure: state.measure,
            command: state.command,
            configuration: state.configuration,
            connector: state.connector,
            cron: state.cron,
            datasource: state.datasource,
            model: state.model,
            notification: state.notification,
            plugin: state.plugin,
            process: state.process,
            snapshot: state.snapshot,
            trigger: state.trigger,
            view: state.trigger,
        };
        callback(resolved);
        return resolved;
    }
}