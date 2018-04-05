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

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	"github.com/yroffin/go-boot-sqllite/core/models"
	app_models "github.com/yroffin/go-jarvis/models"
)

// Property internal members
type Property struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/properties"`
	// Swagger with injection mecanism
	Swagger core_apis.ISwaggerService `@autowired:"swagger"`
}

// IProperty implements IBean
type IProperty interface {
	core_apis.IAPI
}

// New constructor
func (p *Property) New() IProperty {
	bean := Property{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject Property
func (p *Property) SetSwagger(value interface{}) {
	if assertion, ok := value.(core_apis.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Property) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.PropertyBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.PropertyBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Property) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Property) Validate(name string) error {
	return nil
}

// HandlerTasksByID return task by id
func (p *Property) HandlerTasksByID(id string, name string, body string) (interface{}, error) {
	return "", nil
}
