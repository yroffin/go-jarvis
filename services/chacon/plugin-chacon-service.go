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
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	core_services "github.com/yroffin/go-boot-sqllite/core/services"
	app_services "github.com/yroffin/go-jarvis/services"
)

// PluginChaconService internal members
type PluginChaconService struct {
	// members
	*core_services.SERVICE
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginChaconService implements IBean
type IPluginChaconService interface {
	// Extend bean
	core_bean.IBean
	// Local method
	Call(body string) (core_models.IValueBean, error)
}

// New constructor
func (p *PluginChaconService) New() IPluginChaconService {
	bean := PluginChaconService{SERVICE: &core_services.SERVICE{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetPropertyService injection
func (p *PluginChaconService) SetPropertyService(value interface{}) {
	if assertion, ok := value.(app_services.IPropertyService); ok {
		p.PropertyService = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this SERVICE
func (p *PluginChaconService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PluginChaconService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginChaconService) Validate(name string) error {
	return nil
}

// Call execution
func (p *PluginChaconService) Call(body string) (core_models.IValueBean, error) {
	result := (&core_models.ValueBean{}).New()
	return result, nil
}
