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
package shell

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("shell-service", (&ShellService{}).New())
}

// ShellService internal members
type ShellService struct {
	// members
	*winter.Service
	// SetPluginShellService with injection mecanism
	PluginShellService IPluginShellService `@autowired:"plugin-shell-service"`
}

// IShellService implements IBean
type IShellService interface {
	winter.IBean
	AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *ShellService) New() IShellService {
	bean := ShellService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *ShellService) Init() error {
	return nil
}

// PostConstruct this API
func (p *ShellService) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *ShellService) Validate(name string) error {
	return nil
}

// AsObject execution
func (p *ShellService) AsObject(body core_models.IValueBean, args map[string]interface{}) (core_models.IValueBean, error) {
	result, _ := p.PluginShellService.Call(body.GetAsString("body"))
	return result, nil
}

// AsBoolean execution
func (p *ShellService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	return result, nil
}
