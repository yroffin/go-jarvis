// MIT License
//
// Copyright (c) 2017 yroffin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	// Apis
	"flag"

	"github.com/yroffin/go-boot-sqllite/core/apis"
	"github.com/yroffin/go-boot-sqllite/core/business"
	"github.com/yroffin/go-boot-sqllite/core/manager"
	"github.com/yroffin/go-boot-sqllite/core/stores"
	"github.com/yroffin/go-jarvis/apis/commands"
	"github.com/yroffin/go-jarvis/apis/configurations"
	"github.com/yroffin/go-jarvis/apis/connectors"
	"github.com/yroffin/go-jarvis/apis/datasources"
	"github.com/yroffin/go-jarvis/apis/devices"
	"github.com/yroffin/go-jarvis/apis/events"
	"github.com/yroffin/go-jarvis/apis/process"
	"github.com/yroffin/go-jarvis/apis/scripts"
	"github.com/yroffin/go-jarvis/apis/system"
	"github.com/yroffin/go-jarvis/apis/views"
	app_services "github.com/yroffin/go-jarvis/services"
	app_chacon "github.com/yroffin/go-jarvis/services/chacon"
	app_lua "github.com/yroffin/go-jarvis/services/lua"
	app_shell "github.com/yroffin/go-jarvis/services/shell"
	app_slack "github.com/yroffin/go-jarvis/services/slack"
	app_zway "github.com/yroffin/go-jarvis/services/zway"
)

// Rest()
func main() {
	// declare manager and boot it
	m := (&manager.Manager{}).New("manager")
	m.Init()
	// Command Line
	flag.String("Djarvis.slack.api", "", "Slack API")
	m.CommandLine()
	// Core beans
	m.Register("swagger", (&apis.SwaggerService{}).New())
	m.Register("router", (&apis.Router{}).New())
	m.Register("sql-crud-business", (&business.SqlCrudBusiness{}).New())
	m.Register("graph-crud-business", (&business.GraphCrudBusiness{}).New())
	m.Register("sqllite-manager", (&stores.Store{}).New([]string{"CommandBean", "NotificationBean", "SnapshotBean", "ConfigBean", "ScriptPluginBean", "DeviceBean", "ViewBean", "CronBean", "TriggerBean", "ConnectorBean", "PropertyBean", "ProcessusBean"}, "./sqllite.db"))
	m.Register("cayley-manager", (&stores.Graph{}).New([]string{}, "./cayley.db"))
	// helpers
	m.Register("property-service", (&app_services.PropertyService{}).New())
	// PLUGINS beans
	m.Register("plugin-slack-service", (&app_slack.PluginSlackService{}).New())
	m.Register("plugin-shell-service", (&app_shell.PluginShellService{}).New())
	m.Register("plugin-lua-service", (&app_lua.PluginLuaService{}).New())
	m.Register("plugin-chacon-service", (&app_chacon.PluginChaconService{}).New())
	m.Register("plugin-zway-service", (&app_zway.PluginZwayService{}).New())
	// SERVCE beans
	m.Register("slack-service", (&app_slack.SlackService{}).New())
	m.Register("lua-service", (&app_lua.LuaService{}).New())
	m.Register("shell-service", (&app_shell.ShellService{}).New())
	m.Register("zway-service", (&app_zway.ZwayService{}).New())
	m.Register("chacon-service", (&app_chacon.ChaconService{}).New())
	// API beans
	m.Register("CommandBean", (&commands.Command{}).New())
	m.Register("NotificationBean", (&events.Notification{}).New())
	m.Register("SnapshotBean", (&system.Snapshot{}).New())
	m.Register("ConfigBean", (&configurations.Config{}).New())
	m.Register("ScriptPluginBean", (&scripts.ScriptPlugin{}).New())
	m.Register("DeviceBean", (&devices.Device{}).New())
	m.Register("ViewBean", (&views.View{}).New())
	m.Register("CronBean", (&events.Cron{}).New())
	m.Register("TriggerBean", (&events.Trigger{}).New())
	m.Register("ConnectorBean", (&connectors.Connector{}).New())
	m.Register("PropertyBean", (&configurations.Property{}).New())
	m.Register("ProcessBean", (&process.Processus{}).New())
	m.Register("ModelBean", (&process.Model{}).New())
	m.Register("MeasureBean", (&datasources.Measure{}).New())
	m.Register("DataSourceBean", (&datasources.DataSource{}).New())
	m.Register("ConfigurationBean", (&configurations.Configuration{}).New())
	m.Register("SecurityBean", (&system.Security{}).New())
	m.Boot()
}
