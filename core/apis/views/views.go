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
package views

import (
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/apis/devices"
)

func init() {
	winter.Helper.Register("ViewBean", (&View{}).New())
}

// View internal members
type View struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/views"`
	// Device with injection mecanism
	LinkDevice devices.IDevice `@autowired:"DeviceBean" @link:"/api/views" @href:"devices"`
	// Devices
	Device devices.IDevice `@autowired:"DeviceBean"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IView implements IBean
type IView interface {
	engine.IAPI
}

// New constructor
func (p *View) New() IView {
	bean := View{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *View) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ViewBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ViewBeans{}).New()
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
	links, count, err := p.LoadAllLinks("devices",
		func() models.IPersistent {
			return (&devices.DeviceBean{}).New()
		}, winter.Helper.GetBean("DeviceBean").(engine.IAPI))
	return links, count, err
}
