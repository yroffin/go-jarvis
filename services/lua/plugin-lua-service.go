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

	lua "github.com/Shopify/go-lua"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	app_services "github.com/yroffin/go-jarvis/services"
)

func init() {
	engine.Winter.Register("plugin-lua-service", (&PluginLuaService{}).New())
}

// PluginLuaService internal members
type PluginLuaService struct {
	// members
	*engine.SERVICE
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginLuaService implements IBean
type IPluginLuaService interface {
	// Extend bean
	engine.IBean
	// Local method
	Call(body string) (core_models.IValueBean, error)
}

// New constructor
func (p *PluginLuaService) New() IPluginLuaService {
	bean := PluginLuaService{SERVICE: &engine.SERVICE{Bean: &engine.Bean{}}}
	return &bean
}

// SetPropertyService injection
func (p *PluginLuaService) SetPropertyService(value interface{}) {
	if assertion, ok := value.(app_services.IPropertyService); ok {
		p.PropertyService = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this SERVICE
func (p *PluginLuaService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PluginLuaService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginLuaService) Validate(name string) error {
	return nil
}

// Call execution
func (p *PluginLuaService) Call(body string) (core_models.IValueBean, error) {
	l := lua.NewState()
	lua.OpenLibraries(l)

	if err := lua.DoString(l, body); err != nil {
		panic(err)
	}

	result := (&core_models.ValueBean{}).New()
	return result, nil
}
