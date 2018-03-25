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
	core_manager "github.com/yroffin/go-boot-sqllite/core/manager"
	"github.com/yroffin/go-boot-sqllite/core/stores"
	app_apis "github.com/yroffin/go-jarvis/apis"
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
	m := core_manager.Manager{}
	m.Init()
	// Command Line
	flag.String("Djarvis.slack.api", "", "Slack API")
	m.CommandLine()
	// Core beans
	m.Register("swagger", (&apis.SwaggerService{}).New())
	m.Register("router", (&apis.Router{}).New())
	m.Register("sql-crud-business", (&business.SqlCrudBusiness{}).New())
	m.Register("graph-crud-business", (&business.GraphCrudBusiness{}).New())
	m.Register("sqllite-manager", (&stores.Store{}).New([]string{"Node"}, "./sqllite.db"))
	m.Register("cayley-manager", (&stores.Graph{}).New([]string{"Node"}, "./cayley.db"))
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
	m.Register("command-api", (&app_apis.Command{}).New())
	m.Register("notification-api", (&app_apis.Notification{}).New())
	// HREF beans
	//m.Register("command-notification-href", (&app_href.CommandNotificationHref{}).New())
	m.Boot("router")
}
