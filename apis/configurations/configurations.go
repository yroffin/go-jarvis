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
package configurations

import (
	"log"
	"reflect"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
)

func init() {
	engine.Winter.Register("ConfigurationBean", (&Configuration{}).New())
}

// Configuration internal members
type Configuration struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/configurations"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IConfiguration implements IBean
type IConfiguration interface {
	engine.IAPI
}

// New constructor
func (p *Configuration) New() IConfiguration {
	bean := Configuration{API: &engine.API{Bean: &engine.Bean{}}}
	return &bean
}

// SetSwagger inject Configuration
func (p *Configuration) SetSwagger(value interface{}) {
	if assertion, ok := value.(engine.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Configuration) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ConfigurationBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ConfigurationBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Configuration) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Configuration) Validate(name string) error {
	return nil
}

// HandlerTasksByID return task by id
func (p *Configuration) HandlerTasksByID(id string, name string, body string) (interface{}, error) {
	return "", nil
}
