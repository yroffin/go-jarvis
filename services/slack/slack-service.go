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

	"github.com/yroffin/go-boot-sqllite/core/engine"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

func init() {
	engine.Winter.Register("slack-service", (&SlackService{}).New())
}

// SlackService internal members
type SlackService struct {
	// members
	*engine.SERVICE
	// PluginSlackService with injection mecanism
	PluginSlackService IPluginSlackService `@autowired:"plugin-slack-service"`
}

// ISlackService implements IBean
type ISlackService interface {
	engine.IBean
	AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *SlackService) New() ISlackService {
	bean := SlackService{SERVICE: &engine.SERVICE{Bean: &engine.Bean{}}}
	return &bean
}

// SetPluginSlackService injection
func (p *SlackService) SetPluginSlackService(value interface{}) {
	if assertion, ok := value.(IPluginSlackService); ok {
		p.PluginSlackService = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this SERVICE
func (p *SlackService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *SlackService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *SlackService) Validate(name string) error {
	return nil
}

// AsObject execution
func (p *SlackService) AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error) {
	command := make(map[string]interface{})
	log.Println("Args:", args, "Body:", body)
	command["text"] = body.ToString()
	p.PluginSlackService.Call(command)
	return body, nil
}

// AsBoolean execution
func (p *SlackService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	log.Println("Args:", args, "Body:", body, "Not implemented")
	return result, nil
}
