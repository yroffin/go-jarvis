// Package services for common services
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
package lua

import (
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
)

func init() {
	winter.Helper.Register("lua-service", (&service{}).New())
}

// service internal members
type service struct {
	// members
	*winter.Service
	// PluginLuaService with injection mecanism
	PluginLuaService IPluginLuaService `@autowired:"plugin-lua-service"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
}

// ILuaService implements IBean
type ILuaService interface {
	winter.IService
	AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *service) New() ILuaService {
	bean := service{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *service) Init() error {
	return nil
}

// PostConstruct this API
func (p *service) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *service) Validate(name string) error {
	// Notify system ready
	p.MqttService.PublishMostOne("/system/lua", "ready")
	return nil
}

/**
local m = require("extends")

local http = require("http")
response, error = http.request("put", "http://hue/api/xxxxxxxxxxxxxxxxxxxxx/lights/" .. input.light .. "/state", {
    body="{\"on\":" .. input.state .. "}",
    headers={
    }
})

if response.status_code == 200 then
    return { status=response.status_code, output=m.ArrayToJSON(response.body) }
else
    return { status=response.status_code, error=m.ArrayToJSON(response.body) }
end
*/

// AsObject execution
func (p *service) AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error) {
	result, _ := p.PluginLuaService.Call(body.GetAsString("body"), args)
	return result, nil
}

// AsBoolean execution
func (p *service) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	return result, nil
}
