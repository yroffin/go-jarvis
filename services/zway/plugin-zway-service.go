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
package services

import (
	"log"
	"reflect"

	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_services "github.com/yroffin/go-boot-sqllite/core/services"
	app_models "github.com/yroffin/go-jarvis/models"
	app_services "github.com/yroffin/go-jarvis/services"
)

// PluginLuaService internal members
type PluginZwayService struct {
	// members
	*core_services.SERVICE
	// SetPropertyService with injection mecanism
	SetPropertyService func(interface{}) `bean:"property-service"`
	PropertyService    *app_services.PropertyService
}

// IPluginZwayService implements IBean
type IPluginZwayService interface {
	// Extend bean
	core_bean.IBean
	// Local method
	Call(body string) (app_models.IValueBean, error)
}

// New constructor
func (p *PluginZwayService) New() IPluginZwayService {
	bean := PluginZwayService{SERVICE: &core_services.SERVICE{Bean: &core_bean.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginZwayService) Init() error {
	// inject store
	p.SetPropertyService = func(value interface{}) {
		if assertion, ok := value.(*app_services.PropertyService); ok {
			p.PropertyService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	return nil
}

// PostConstruct this SERVICE
func (p *PluginZwayService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginZwayService) Validate(name string) error {
	return nil
}

// Call execution
func (p *PluginZwayService) Call(body string) (app_models.IValueBean, error) {
	result := (&app_models.ValueBean{}).New()
	return result, nil
}
