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
	app_helpers "github.com/yroffin/go-jarvis/helpers"
	app_services "github.com/yroffin/go-jarvis/services"
)

// PluginSlackService internal members
type PluginSlackService struct {
	// members
	*core_services.SERVICE
	// SetPropertyService with injection mecanism
	SetPropertyService func(interface{}) `bean:"property-service"`
	PropertyService    *app_services.PropertyService
}

// IPluginSlackService implements IBean
type IPluginSlackService interface {
	core_bean.IBean
}

// New constructor
func (p *PluginSlackService) New() IPluginSlackService {
	bean := PluginSlackService{SERVICE: &core_services.SERVICE{Bean: &core_bean.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginSlackService) Init() error {
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
func (p *PluginSlackService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginSlackService) Validate(name string) error {
	return nil
}

// Call execution
func (p *PluginSlackService) Call(body map[string]interface{}) (map[string]interface{}, error) {
	client := &app_helpers.HTTPClient{URL: p.PropertyService.Get("jarvis.slack.url", "https://hooks.slask.com/services")}
	headers := make(map[string]string)
	params := make(map[string]string)
	result, err := client.POST(p.PropertyService.Get("jarvis.slack.api", ""), body, headers, params)
	return result, err
}
