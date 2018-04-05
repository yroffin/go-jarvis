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
	"log"
	"reflect"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-jarvis/apis/events"
	"github.com/yroffin/go-jarvis/apis/scripts"
	app_models "github.com/yroffin/go-jarvis/models"
)

// Device internal members
type Device struct {
	// Base component
	*core_apis.API
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
	Swagger core_apis.ISwaggerService `@autowired:"swagger"`
}

// IDevice implements IBean
type IDevice interface {
	core_apis.IAPI
}

// New constructor
func (p *Device) New() IDevice {
	bean := Device{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject Device
func (p *Device) SetSwagger(value interface{}) {
	if assertion, ok := value.(core_apis.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetDevice inject notification
func (p *Device) SetDevice(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.Device = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetLinkDevice injection
func (p *Device) SetLinkDevice(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.LinkDevice = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetTrigger inject notification
func (p *Device) SetTrigger(value interface{}) {
	if assertion, ok := value.(events.ITrigger); ok {
		p.Trigger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetLinkTrigger injection
func (p *Device) SetLinkTrigger(value interface{}) {
	if assertion, ok := value.(events.ITrigger); ok {
		p.LinkTrigger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetPluginScript inject notification
func (p *Device) SetPluginScript(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.PluginScript = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetLinkPluginScript injection
func (p *Device) SetLinkPluginScript(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.LinkPluginScript = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Device) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.DeviceBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.DeviceBeans{}).New()
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
