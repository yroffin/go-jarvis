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
package slack

import (
	"github.com/yroffin/go-boot-sqllite/core/winter"
	app_helpers "github.com/yroffin/go-jarvis/core/helpers"
	app_services "github.com/yroffin/go-jarvis/core/services"
)

func init() {
	winter.Helper.Register("plugin-slack-service", (&PluginSlackService{}).New())
}

// PluginSlackService internal members
type PluginSlackService struct {
	// members
	*winter.Service
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginSlackService implements IBean
type IPluginSlackService interface {
	winter.IBean
	Call(body map[string]interface{}) (map[string]interface{}, error)
}

// New constructor
func (p *PluginSlackService) New() IPluginSlackService {
	bean := PluginSlackService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginSlackService) Init() error {
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
