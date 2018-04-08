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
package zway

import (
	"log"
	"strings"

	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("zway-service", (&ZwayService{}).New())
}

// ZwayService internal members
type ZwayService struct {
	// members
	*winter.Service
	// SetPluginZwayService with injection mecanism
	PluginZwayService IPluginZwayService `@autowired:"plugin-zway-service"`
}

// IZwayService implements IBean
type IZwayService interface {
	winter.IBean
	AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error)
	AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error)
}

// New constructor
func (p *ZwayService) New() IZwayService {
	bean := ZwayService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
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
func (p *ZwayService) AsObject(body models.IValueBean, args map[string]interface{}) (models.IValueBean, error) {
	res := strings.Split(strings.Split(body.GetAsString("body"), "_")[2], "-")
	result, _ := p.PluginZwayService.Call(res[0], res[1], res[2], res[3])
	return result, nil
}

// AsBoolean execution
func (p *ZwayService) AsBoolean(body map[string]interface{}, args map[string]interface{}) (bool, error) {
	result := false
	log.Println("Not implemented")
	return result, nil
}
