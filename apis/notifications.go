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

// Bean internal members
type Notification struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	Crud map[string]interface{} `@crud:"/api/notifications"`
	// SwaggerService with injection mecanism
	SwaggerService core_apis.ISwaggerService `@autowired:"swagger"`
}

// INotification implements IBean
type INotification interface {
	core_bean.IBean
}

// New constructor
func (p *Notification) New() INotification {
	bean := Notification{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject notification
func (p *Notification) SetSwagger(value interface{}) {
	if assertion, ok := value.(*core_apis.SwaggerService); ok {
		p.SwaggerService = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Notification) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.NotificationBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.NotificationBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Notification) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.SwaggerService, p)
	return nil
}

// Validate this API
func (p *Notification) Validate(name string) error {
	return nil
}
