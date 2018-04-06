// Package interfaces for common interfaces
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
package devices

import (
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/apis/events"
	"github.com/yroffin/go-jarvis/apis/scripts"
)

func init() {
	winter.Helper.Register("DeviceBean", (&Device{}).New())
}

// Device internal members
type Device struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/devices"`
	// Device with injection mecanism
	LinkDevice IDevice `@autowired:"DeviceBean" @link:"/api/devices" @href:"devices"`
	Device     IDevice `@autowired:"DeviceBean"`
	// Trigger with injection mecanism
	LinkTrigger events.ITrigger `@autowired:"TriggerBean" @link:"/api/devices" @href:"triggers"`
	Trigger     events.ITrigger `@autowired:"TriggerBean"`
	// PluginScript with injection mecanism
	LinkPluginScript scripts.IScriptPlugin `@autowired:"ScriptPluginBean" @link:"/api/devices" @href:"plugins/scripts"`
	PluginScript     scripts.IScriptPlugin `@autowired:"ScriptPluginBean"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IDevice implements IBean
type IDevice interface {
	engine.IAPI
}

// New constructor
func (p *Device) New() IDevice {
	bean := Device{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Device) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&DeviceBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&DeviceBeans{}).New()
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		if name == "uml" {
			// task
			return p.Uml(id, body)
		}
		return "", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Device) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Device) Validate(name string) error {
	return nil
}

// HandlerTasksByID return task by id
func (p *Device) Uml(id string, body string) (interface{}, int, error) {
	return "", -1, nil
}
