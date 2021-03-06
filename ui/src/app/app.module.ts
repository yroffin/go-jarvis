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

import { environment } from '../environments/environment';
import { BrowserModule } from '@angular/platform-browser';
import { CommonModule } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { StoreModule } from '@ngrx/store';

import { AppComponent } from './app.component';

import { HttpClientModule } from '@angular/common/http';

/**
 * material2
 */
import { MatSidenavModule, MatIcon, MatToolbar, MatDividerModule, MatChipsModule } from '@angular/material';
import { MatCheckboxModule } from '@angular/material';
import { MatButtonModule } from '@angular/material';
import { MatGridListModule } from '@angular/material';
import { MatInputModule } from '@angular/material';
import { MatTableModule } from '@angular/material';
import { MatTabsModule } from '@angular/material';
import { MatSelectModule } from '@angular/material';
import { MatOptionModule } from '@angular/material';
import { MatCardModule } from '@angular/material';
import { MatSnackBarModule } from '@angular/material';
import { MatFormFieldModule } from '@angular/material';
import { MatPaginatorModule } from '@angular/material';
import { MatIconModule, MatIconRegistry } from '@angular/material';
import { MatBadgeModule } from '@angular/material/badge';
import { MatExpansionModule } from '@angular/material';
import { MatMenuModule } from '@angular/material';
import { MatToolbarModule } from '@angular/material';
import { MatDialogModule } from '@angular/material/dialog';

/**
 * primeng
 */
import { ButtonModule } from 'primeng/primeng';
import { ChartModule } from 'primeng/primeng';
import { DataTableModule, SharedModule } from 'primeng/primeng';
import { MenubarModule, MenuModule } from 'primeng/primeng';
import { CheckboxModule } from 'primeng/primeng';
import { InputTextModule } from 'primeng/primeng';
import { AccordionModule } from 'primeng/primeng';
import { CodeHighlighterModule } from 'primeng/primeng';
import { InputTextareaModule } from 'primeng/primeng';
import { DataListModule } from 'primeng/primeng';
import { TabViewModule } from 'primeng/primeng';
import { DataGridModule } from 'primeng/primeng';
import { PanelModule } from 'primeng/primeng';
import { GrowlModule } from 'primeng/primeng';
import { MessagesModule } from 'primeng/primeng';
import { StepsModule } from 'primeng/primeng';
import { PanelMenuModule } from 'primeng/primeng';
import { DialogModule } from 'primeng/primeng';
import { FieldsetModule } from 'primeng/primeng';
import { DropdownModule } from 'primeng/primeng';
import { ConfirmDialogModule, ConfirmationService } from 'primeng/primeng';
import { SplitButtonModule } from 'primeng/primeng';
import { ToolbarModule } from 'primeng/primeng';
import { TooltipModule } from 'primeng/primeng';
import { TreeTableModule } from 'primeng/primeng';
import { CalendarModule } from 'primeng/primeng';
import { SpinnerModule } from 'primeng/primeng';
import { SliderModule } from 'primeng/primeng';
import { ToggleButtonModule } from 'primeng/primeng';
import { SidebarModule } from 'primeng/sidebar';
import { OverlayPanelModule } from 'primeng/overlaypanel';
import { ScrollPanelModule } from 'primeng/scrollpanel';
import { ContextMenuModule } from 'primeng/contextmenu';

import { WindowRef } from './service/jarvis-utils.service';
import { LoggerService } from './service/logger.service';
import { NavigationGuard } from './guard/navigation.service';
import { ProfileGuard } from './guard/profile.service';

import { JarvisMqttService } from './service/jarvis-mqtt.service';
import { JarvisSecurityService } from './service/jarvis-security.service';
import { JarvisConfigurationService } from './service/jarvis-configuration.service';
import { JarvisDataDeviceService } from './service/jarvis-data-device.service';
import { JarvisDataTriggerService } from './service/jarvis-data-trigger.service';
import { JarvisDataPluginService } from './service/jarvis-data-plugin.service';
import { JarvisDataCommandService } from './service/jarvis-data-command.service';
import { JarvisDataConfigurationService } from './service/jarvis-data-configuration.service';
import { JarvisDataPropertyService } from './service/jarvis-data-property.service';
import { JarvisDataNotificationService } from './service/jarvis-data-notification.service';
import { JarvisDataCronService } from './service/jarvis-data-cron.service';
import { JarvisDataProcessService } from './service/jarvis-data-process.service';
import { JarvisDataSnapshotService } from './service/jarvis-data-snapshot.service';
import { JarvisDataViewService } from './service/jarvis-data-view.service';
import { JarvisDataRawService } from './service/jarvis-data-raw.service';
import { JarvisDataDatasourceService } from './service/jarvis-data-datasource.service';
import { JarvisDataModelService } from './service/jarvis-data-model.service';
import { JarvisLoaderService } from './service/jarvis-loader.service';
import { JarvisMessageService } from './service/jarvis-message.service';

import { JarvisLayoutDirective } from './directive/jarvis-layout.directive';

import { JarvisHomeComponent } from './component/jarvis-home/jarvis-home.component';
import { JarvisResourcesComponent } from './component/jarvis-resources/jarvis-resources.component';
import { JarvisResourceDeviceComponent } from './component/jarvis-resource-device/jarvis-resource-device.component';
import { JarvisResourcePluginComponent } from './component/jarvis-resource-plugin/jarvis-resource-plugin.component';
import { JarvisResourceCommandComponent } from './component/jarvis-resource-command/jarvis-resource-command.component';
import { JarvisResourceTriggerComponent } from './component/jarvis-resource-trigger/jarvis-resource-trigger.component';
import { JarvisResourceCronComponent } from './component/jarvis-resource-cron/jarvis-resource-cron.component';
import { JarvisResourceConfigurationComponent } from './component/jarvis-resource-configuration/jarvis-resource-configuration.component';
import { JarvisResourceNotificationComponent } from './component/jarvis-resource-notification/jarvis-resource-notification.component';
import { JarvisResourcePropertyComponent } from './component/jarvis-resource-property/jarvis-resource-property.component';
import { JarvisResourceConnectorComponent } from './component/jarvis-resource-connector/jarvis-resource-connector.component'
import { JarvisResourceViewComponent } from './component/jarvis-resource-view/jarvis-resource-view.component'
import { JarvisInlineSvgDirective } from './directive/jarvis-inline-svg.directive';

import { JarvisTileComponent } from './component/jarvis-tile/jarvis-tile.component';
import { JarvisToolbarResourceComponent } from './component/jarvis-toolbar-resource/jarvis-toolbar-resource.component';
import { JarvisLoginComponent } from './component/jarvis-login/jarvis-login.component';
import { JarvisResourceSnapshotComponent } from './component/jarvis-resource-snapshot/jarvis-resource-snapshot.component';
import { JarvisDesktopComponent, OrderSortPipe } from './component/jarvis-desktop/jarvis-desktop.component';
import { JarvisResourceDatasourceComponent } from './component/jarvis-resource-datasource/jarvis-resource-datasource.component';

/**
 * stores
 */
import { BrokerStoreService } from './store/broker.store';
import { MessageStoreService } from './store/message.store';
import { ViewStoreService } from './store/view.store';

import { JarvisServerResourcesComponent } from './component/jarvis-server-resources/jarvis-server-resources.component';
import { JarvisBrokerComponent } from './component/jarvis-broker/jarvis-broker.component';
import { JarvisResourceProcessComponent } from './component/jarvis-resource-process/jarvis-resource-process.component';
import { JarvisResourceModelComponent } from './component/jarvis-resource-model/jarvis-resource-model.component';
import { JarvisGraphComponent } from './widget/jarvis-graph/jarvis-graph.component';
import { SystemStoreService } from './store/system.store';
import { JarvisGraphBrowserComponent } from './component/jarvis-graph-browser/jarvis-graph-browser.component';
import { GraphStoreService } from './store/graph.store';
import { JarvisBpmnComponent } from './widget/jarvis-bpmnjs/jarvis-bpmnjs.component';
import { JarvisGraphExplorerComponent } from './widget/jarvis-graph-explorer/jarvis-graph-explorer.component';
import { DialogConfirmDrop } from './dialog/drop-resource/jarvis-drop-resource.component';
import { DialogAbout } from './dialog/about/jarvis-about.component';
import { DialogPickResource } from './dialog/picker/jarvis-pick-resource.component';
import { DeviceResolver } from './resolver/device-resolver';
import { ResourceStoreService } from './store/resources.store';
import { ViewResolver } from './resolver/view-resolver';
import { PluginResolver } from './resolver/plugin-resolver';
import { CommandResolver } from './resolver/command-resolver';
import { ConnectorResolver } from './resolver/connector-resolver';
import { TriggerResolver } from './resolver/trigger-resolver';
import { CronResolver } from './resolver/cron-resolver';
import { NotificationResolver } from './resolver/notification-resolver';
import { ModelResolver } from './resolver/model-resolver';
import { ProcessResolver } from './resolver/process-resolver';
import { ConfigurationResolver } from './resolver/configuration-resolver';
import { PropertyResolver } from './resolver/property-resolver';
import { SnapshotResolver } from './resolver/snapshot-resolver';
import { DatasourceResolver } from './resolver/datasource-resolver';

/**
 * default route definition
 */
const appRoutes: Routes = [
  { path: 'devices', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'devices' } },
  { path: 'devices/:id', component: JarvisResourceDeviceComponent, canActivate: [ProfileGuard], resolve: { device: DeviceResolver } },
  { path: 'plugins', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'plugins' } },
  { path: 'plugins/:id', component: JarvisResourcePluginComponent, canActivate: [ProfileGuard], resolve: { plugin: PluginResolver } },
  { path: 'commands', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'commands' } },
  { path: 'commands/:id', component: JarvisResourceCommandComponent, canActivate: [ProfileGuard], resolve: { command: CommandResolver } },
  { path: 'triggers', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'triggers' } },
  { path: 'triggers/:id', component: JarvisResourceTriggerComponent, canActivate: [ProfileGuard], resolve: { trigger: TriggerResolver } },
  { path: 'crons', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'crons' } },
  { path: 'crons/:id', component: JarvisResourceCronComponent, canActivate: [ProfileGuard], resolve: { cron: CronResolver } },
  { path: 'processes', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'processes' } },
  { path: 'processes/:id', component: JarvisResourceProcessComponent, canActivate: [ProfileGuard], resolve: { process: ProcessResolver } },
  { path: 'configurations', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'configurations' } },
  { path: 'configurations/:id', component: JarvisResourceConfigurationComponent, canActivate: [ProfileGuard], resolve: { configuration: ConfigurationResolver } },
  { path: 'notifications', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'notifications' } },
  { path: 'notifications/:id', component: JarvisResourceNotificationComponent, canActivate: [ProfileGuard], resolve: { notification: NotificationResolver } },
  { path: 'properties', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'properties' } },
  { path: 'properties/:id', component: JarvisResourcePropertyComponent, canActivate: [ProfileGuard], resolve: { property: PropertyResolver } },
  { path: 'connectors', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'connectors' } },
  { path: 'connectors/:id', component: JarvisResourceConnectorComponent, canActivate: [ProfileGuard], resolve: { command: ConnectorResolver } },
  { path: 'views', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'views' } },
  { path: 'views/:id', component: JarvisResourceViewComponent, canActivate: [ProfileGuard], resolve: { view: ViewResolver } },
  { path: 'snapshots', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'snapshots' } },
  { path: 'snapshots/:id', component: JarvisResourceSnapshotComponent, canActivate: [ProfileGuard], resolve: { snapshot: SnapshotResolver } },
  { path: 'datasources', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'datasources' } },
  { path: 'datasources/:id', component: JarvisResourceDatasourceComponent, canActivate: [ProfileGuard], resolve: { datasource: DatasourceResolver } },
  { path: 'models', component: JarvisResourcesComponent, canActivate: [ProfileGuard, NavigationGuard], data: { resource: 'models' } },
  { path: 'models/:id', component: JarvisResourceModelComponent, canActivate: [ProfileGuard, NavigationGuard], resolve: { model: ModelResolver } },
  { path: 'graph', component: JarvisGraphBrowserComponent, canActivate: [ProfileGuard, NavigationGuard] },
  { path: 'resources', component: JarvisServerResourcesComponent, canActivate: [ProfileGuard, NavigationGuard] },
  { path: 'broker', component: JarvisBrokerComponent, canActivate: [ProfileGuard, NavigationGuard] },
  { path: 'desktop', component: JarvisDesktopComponent, canActivate: [ProfileGuard] },
  { path: 'login', component: JarvisLoginComponent, canActivate: [NavigationGuard] },
  { path: '', component: JarvisDesktopComponent, canActivate: [ProfileGuard] },
  { path: '**', component: JarvisDesktopComponent, canActivate: [ProfileGuard] }
];

@NgModule({
  declarations: [
    AppComponent,
    JarvisLayoutDirective,
    JarvisHomeComponent,
    JarvisResourceDeviceComponent,
    JarvisResourcesComponent,
    JarvisInlineSvgDirective,
    JarvisTileComponent,
    JarvisToolbarResourceComponent,
    JarvisResourcePluginComponent,
    JarvisResourceCommandComponent,
    JarvisResourceTriggerComponent,
    JarvisResourceConnectorComponent,
    JarvisResourceCronComponent,
    JarvisResourceConfigurationComponent,
    JarvisResourceNotificationComponent,
    JarvisResourcePropertyComponent,
    JarvisResourceViewComponent,
    JarvisLoginComponent,
    JarvisResourceSnapshotComponent,
    JarvisDesktopComponent,
    JarvisResourceDatasourceComponent,
    JarvisServerResourcesComponent,
    JarvisBrokerComponent,
    JarvisResourceProcessComponent,
    JarvisResourceModelComponent,
    JarvisGraphBrowserComponent,
    JarvisGraphComponent,
    JarvisBpmnComponent,
    JarvisGraphExplorerComponent,
    DialogConfirmDrop,
    DialogAbout,
    DialogPickResource,
    OrderSortPipe
  ],
  entryComponents: [
    DialogConfirmDrop,
    DialogAbout,
    DialogPickResource
  ],
  imports: [
    CommonModule,
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
    /**
     * material2
     */
    MatSidenavModule,
    MatButtonModule,
    MatGridListModule,
    MatInputModule,
    MatCheckboxModule,
    MatTableModule,
    MatTabsModule,
    MatOptionModule,
    MatSelectModule,
    MatCardModule,
    MatSnackBarModule,
    MatFormFieldModule,
    MatIconModule,
    MatPaginatorModule,
    MatBadgeModule,
    MatExpansionModule,
    MatMenuModule,
    MatToolbarModule,
    MatDialogModule,
    MatDividerModule,
    MatChipsModule,
    /**
     * primeface
     */
    DataTableModule,
    SharedModule,
    MenuModule,
    MenubarModule,
    CheckboxModule,
    InputTextModule,
    AccordionModule,
    CodeHighlighterModule,
    InputTextareaModule,
    DataListModule,
    TabViewModule,
    DataGridModule,
    PanelModule,
    GrowlModule,
    MessagesModule,
    StepsModule,
    ButtonModule,
    PanelMenuModule,
    DialogModule,
    FieldsetModule,
    DropdownModule,
    ConfirmDialogModule,
    SplitButtonModule,
    ToolbarModule,
    TooltipModule,
    TreeTableModule,
    ChartModule,
    CalendarModule,
    SpinnerModule,
    SliderModule,
    ToggleButtonModule,
    SidebarModule,
    OverlayPanelModule,
    ScrollPanelModule,
    ContextMenuModule,
    /**
     * routes
     */
    RouterModule.forRoot(appRoutes, { enableTracing: environment.enableTracing }),
    /**
     * store
     */
    StoreModule.forRoot({
      broker: BrokerStoreService.reducer,
      message: MessageStoreService.reducer,
      view: ViewStoreService.reducer,
      system: SystemStoreService.reducer,
      graph: GraphStoreService.reducer,
      resources: ResourceStoreService.reducer
    })
  ],
  providers: [
    /**
     * extends
     */
    WindowRef,
    MatIconRegistry,
    /**
     * jarvis
     */
    JarvisConfigurationService,
    JarvisSecurityService,
    JarvisDataDeviceService,
    JarvisDataTriggerService,
    JarvisDataPluginService,
    JarvisDataCommandService,
    JarvisDataConfigurationService,
    JarvisDataPropertyService,
    JarvisDataCronService,
    JarvisDataNotificationService,
    JarvisDataProcessService,
    JarvisDataRawService,
    JarvisDataViewService,
    JarvisDataSnapshotService,
    JarvisDataViewService,
    JarvisDataDatasourceService,
    JarvisDataModelService,
    JarvisMqttService,
    /**
     * guards
     */
    NavigationGuard,
    ProfileGuard,
    LoggerService,
    JarvisLoaderService,
    JarvisMessageService,
    ConfirmationService,
    /**
     * store
     */
    BrokerStoreService,
    MessageStoreService,
    ViewStoreService,
    SystemStoreService,
    GraphStoreService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(
    public matIconRegistry: MatIconRegistry) {
    matIconRegistry.registerFontClassAlias('fontawesome', 'fa');
  }
}
