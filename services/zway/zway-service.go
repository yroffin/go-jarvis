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
package services

import (
	"log"
	"reflect"

	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_services "github.com/yroffin/go-boot-sqllite/core/services"
	app_models "github.com/yroffin/go-jarvis/models"
)

// ZwayService internal members
type ZwayService struct {
	// members
	*core_services.SERVICE
	// SetPluginZwayService with injection mecanism
	SetPluginZwayService func(interface{}) `bean:"plugin-zway-service"`
	PluginZwayService    *PluginZwayService
}

// IZwayService implements IBean
type IZwayService interface {
	core_bean.IBean
}

// New constructor
func (p *ZwayService) New() IZwayService {
	bean := ZwayService{SERVICE: &core_services.SERVICE{Bean: &core_bean.Bean{}}}
	return &bean
}

// Init this API
func (p *ZwayService) Init() error {
	// inject store
	p.SetPluginZwayService = func(value interface{}) {
		if assertion, ok := value.(*PluginZwayService); ok {
			p.PluginZwayService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	return nil
}

// PostConstruct this API
func (p *ZwayService) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *ZwayService) Validate(name string) error {
	return nil
}

// AsObject execution
func (p *ZwayService) AsObject(body app_models.IValueBean, args map[string]interface{}) (app_models.IValueBean, error) {
	log.Println("Args:", args, "Body:", body)
	result, _ := p.PluginZwayService.Call(body.GetAsString("body"))
	return result, nil
}

// AsBoolean execution
func (p *ZwayService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	log.Println("Args:", args, "Body:", body, "Not implemented")
	return result, nil
}
