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
	"github.com/yroffin/go-boot-sqllite/core/business"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/stores"
	app_models "github.com/yroffin/go-jarvis/models"
)

// Config internal members
type Config struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/configs"`
	// Swagger with injection mecanism
	Swagger core_apis.ISwaggerService `@autowired:"swagger"`
	// SqlCrudBusiness with injection mecanism
	SQLCrudBusiness business.ICrudBusiness `@autowired:"sql-crud-business"`
	// GraphBusiness with injection mecanism
	GraphBusiness business.ILinkBusiness `@autowired:"graph-crud-business"`
}

// IConfig implements IBean
type IConfig interface {
	core_apis.IAPI
}

// New constructor
func (p *Config) New() IConfig {
	bean := Config{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject Config
func (p *Config) SetSwagger(value interface{}) {
	if assertion, ok := value.(core_apis.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetSQLCrudBusiness inject CrudBusiness
func (p *Config) SetSQLCrudBusiness(value interface{}) {
	if assertion, ok := value.(business.ICrudBusiness); ok {
		p.SQLCrudBusiness = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetGraphBusiness inject CrudBusiness
func (p *Config) SetGraphBusiness(value interface{}) {
	if assertion, ok := value.(business.ILinkBusiness); ok {
		p.GraphBusiness = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Config) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.ConfigBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.ConfigBeans{}).New()
	}
	p.HandlerTasks = func(name string, body string) (interface{}, int, error) {
		if name == "statistics" {
			// task
			return p.Statistics(body)
		}
		return "", -1, nil
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		return "", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Config) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Config) Validate(name string) error {
	return nil
}

// StatisticsResult result
type StatisticsResult struct {
	SQL   []stores.IStats
	Graph []stores.IStats
}

// Statistics this Snapshot
func (p *Config) Statistics(body string) (interface{}, int, error) {
	sql, _ := p.SQLCrudBusiness.Statistics()
	graph, _ := p.GraphBusiness.Statistics()
	return StatisticsResult{SQL: sql, Graph: graph}, -1, nil
}
