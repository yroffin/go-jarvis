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
)

// LuaService internal members
type LuaService struct {
	// members
	*core_services.SERVICE
	// PluginLuaService with injection mecanism
	PluginLuaService IPluginLuaService `@autowired:"plugin-lua-service"`
}

// ILuaService implements IBean
type ILuaService interface {
	core_bean.IBean
	AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *LuaService) New() ILuaService {
	bean := LuaService{SERVICE: &core_services.SERVICE{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetPluginLuaService injection
func (p *LuaService) SetPluginLuaService(value interface{}) {
	if assertion, ok := value.(IPluginLuaService); ok {
		p.PluginLuaService = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *LuaService) Init() error {
	return nil
}

// PostConstruct this API
func (p *LuaService) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *LuaService) Validate(name string) error {
	return nil
}

// AsObject execution
func (p *LuaService) AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error) {
	log.Println("Args:", args, "Body:", body)
	result, _ := p.PluginLuaService.Call(body.GetAsString("body"))
	return result, nil
}

// AsBoolean execution
func (p *LuaService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	log.Println("Args:", args, "Body:", body, "Not implemented")
	return result, nil
}
