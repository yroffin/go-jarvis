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
package apis

import (
	"log"
	"reflect"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	"github.com/yroffin/go-boot-sqllite/core/models"
	app_models "github.com/yroffin/go-jarvis/models"
)

// View internal members
type View struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/views"`
	// Notification with injection mecanism
	LinkDevice IDevice `@autowired:"DeviceBean" @link:"/api/views" @href:"devices"`
	Device     IDevice `@autowired:"DeviceBean"`
	// Swagger with injection mecanism
	Swagger core_apis.ISwaggerService `@autowired:"swagger"`
}

// IView implements IBean
type IView interface {
	core_apis.IAPI
}

// New constructor
func (p *View) New() IView {
	bean := View{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject View
func (p *View) SetSwagger(value interface{}) {
	if assertion, ok := value.(core_apis.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetDevice inject notification
func (p *View) SetDevice(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.Device = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetLinkDevice injection
func (p *View) SetLinkDevice(value interface{}) {
	if assertion, ok := value.(IDevice); ok {
		p.LinkDevice = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *View) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.ViewBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.ViewBeans{}).New()
	}
	p.HandlerTasks = func(name string, body string) (interface{}, int, error) {
		if name == "GET" {
			// task
			return p.GetAllViews(body)
		}
		return "", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *View) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *View) Validate(name string) error {
	return nil
}

// GetAllViews read all views and all data
func (p *View) GetAllViews(body string) (interface{}, int, error) {
	views := (&app_models.ViewBeans{}).New()
	p.SQLCrudBusiness.GetAll(p.Factory(), views)
	for _, view := range views.Get() {
		devices := make([]app_models.DeviceBean, 0)
		edges, _ := p.GraphBusiness.GetAllLink(view.GetID(), make([]models.IEdgeBean, 0))
		for _, edge := range edges {
			device := (&app_models.DeviceBean{}).New()
			device.SetID(edge.GetTargetID())
			p.SQLCrudBusiness.Get(device)
			devices = append(devices, *device.Get())
		}
		view.(app_models.IViewBean).SetDevices(devices)
	}
	return views.Get(), -1, nil
}
