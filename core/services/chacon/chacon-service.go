// Package chacon for common interfaces
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
package chacon

import (
	"strings"

	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
)

func init() {
	winter.Helper.Register("chacon-service", (&service{}).New())
}

// service internal members
type service struct {
	// members
	*winter.Service
	// SetPluginSlackService with injection mecanism
	PluginRFLinkService IPluginRFLinkService `@autowired:"plugin-rflink-service"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
}

// IChaconService implements IBean
type IChaconService interface {
	winter.IService
	AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *service) New() IChaconService {
	bean := service{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *service) Init() error {
	return nil
}

// PostConstruct this API
func (p *service) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *service) Validate(name string) error {
	// Notify system ready
	p.MqttService.PublishMostOne("/system/chacon", "ready")
	return nil
}

func parse(args map[string]interface{}, list []string) []string {
	result := make([]string, 0)
	for _, value := range list {
		for k, v := range args {
			value = strings.Replace(value, "${"+k+"}", v.(string), -1)
		}
		result = append(result, value)
	}
	return result
}

// AsObject execution
func (p *service) AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error) {
	parsed := parse(args, strings.Split(body.GetAsString("body"), " "))
	if len(parsed) == 4 && parsed[0] == "CHACON" && !strings.Contains(parsed[1], "$") && !strings.Contains(parsed[2], "$") && !strings.Contains(parsed[3], "$") {
		result, _ := p.PluginRFLinkService.Chacon(parsed[1], parsed[2], parsed[3])
		return result, nil
	}
	return nil, nil
}

// AsBoolean execution
func (p *service) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	return result, nil
}
