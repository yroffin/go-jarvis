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

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_business "github.com/yroffin/go-boot-sqllite/core/business"
	core_manager "github.com/yroffin/go-boot-sqllite/core/manager"
	core_stores "github.com/yroffin/go-boot-sqllite/core/stores"
	app_apis "github.com/yroffin/go-jarvis/apis"
	app_services "github.com/yroffin/go-jarvis/services"
)

// Rest()
func main() {
	// declare manager and boot it
	var m = core_manager.Manager{}
	m.Init()
	// Command Line
	flag.String("Djarvis.slack.api", "", "Slack API")
	m.CommandLine()
	// Core beans
	m.Register("router", &core_apis.Router{Bean: &core_bean.Bean{}})
	m.Register("crud-business", &core_business.CrudBusiness{Bean: &core_bean.Bean{}})
	m.Register("store-manager", &core_stores.Store{Bean: &core_bean.Bean{}, Tables: []string{"Command", "Notification"}, DbPath: "./database.db"})
	// helpers
	m.Register("property-service", (&app_services.PropertyService{}).New())
	// PLUGINS beans
	m.Register("plugin-slack-service", (&app_services.PluginSlackService{}).New())
	m.Register("plugin-shell-service", (&app_services.PluginShellService{}).New())
	m.Register("plugin-lua-service", (&app_services.PluginLuaService{}).New())
	// SERVCE beans
	m.Register("slack-service", (&app_services.SlackService{}).New())
	m.Register("lua-service", (&app_services.LuaService{}).New())
	m.Register("shell-service", (&app_services.ShellService{}).New())
	m.Register("zway-service", (&app_services.ZwayService{}).New())
	m.Register("chacon-service", (&app_services.ChaconService{}).New())
	// API beans
	m.Register("command", (&app_apis.Command{}).New())
	m.Register("notification", (&app_apis.Notification{}).New())
	m.Boot()
}
