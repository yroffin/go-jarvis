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

	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_services "github.com/yroffin/go-boot-sqllite/core/services"
)

// ZwayService internal members
type ZwayService struct {
	// members
	*core_services.SERVICE
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

// Execute this command
func (p *ZwayService) Execute(id string, body string) (string, error) {
	log.Println("Script:", body)
	return body, nil
}
