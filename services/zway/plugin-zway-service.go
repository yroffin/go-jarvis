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
package zway

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	app_services "github.com/yroffin/go-jarvis/services"
)

func init() {
	winter.Helper.Register("plugin-zway-service", (&PluginZwayService{}).New())
}

// PluginLuaService internal members
type PluginZwayService struct {
	// members
	*winter.Service
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginZwayService implements IBean
type IPluginZwayService interface {
	// Extend bean
	winter.IBean
	// Local method
	Call(body string) (core_models.IValueBean, error)
}

// New constructor
func (p *PluginZwayService) New() IPluginZwayService {
	bean := PluginZwayService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginZwayService) Init() error {
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
func (p *PluginZwayService) Call(body string) (core_models.IValueBean, error) {
	result := (&core_models.ValueBean{}).New()
	return result, nil
}
