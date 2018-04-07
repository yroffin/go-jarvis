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
package slack

import (
	"log"

	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("slack-service", (&SlackService{}).New())
}

// SlackService internal members
type SlackService struct {
	// members
	*winter.Service
	// PluginSlackService with injection mecanism
	PluginSlackService IPluginSlackService `@autowired:"plugin-slack-service"`
}

// ISlackService implements IBean
type ISlackService interface {
	winter.IBean
	AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *SlackService) New() ISlackService {
	bean := SlackService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
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
func (p *SlackService) AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error) {
	command := make(map[string]interface{})
	log.Println("Args:", args, "Body:", body)
	command["text"] = models.ToString(body)
	p.PluginSlackService.Call(command)
	return body, nil
}

// AsBoolean execution
func (p *SlackService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	log.Println("Args:", args, "Body:", body, "Not implemented")
	return result, nil
}
