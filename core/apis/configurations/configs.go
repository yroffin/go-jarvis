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
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("ConfigBean", (&Config{}).New())
}

// Config internal members
type Config struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/configs"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// SqlCrudBusiness with injection mecanism
	SQLCrudBusiness engine.ICrudBusiness `@autowired:"sql-crud-business"`
	// GraphBusiness with injection mecanism
	GraphBusiness engine.ILinkBusiness `@autowired:"graph-crud-business"`
}

// IConfig implements IBean
type IConfig interface {
	engine.IAPI
}

// New constructor
func (p *Config) New() IConfig {
	bean := Config{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Config) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ConfigBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ConfigBeans{}).New()
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
	SQL   []engine.IStats
	Graph []engine.IStats
}

// Statistics this Snapshot
func (p *Config) Statistics(body string) (interface{}, int, error) {
	sql, _ := p.SQLCrudBusiness.Statistics()
	graph, _ := p.GraphBusiness.Statistics()
	return StatisticsResult{SQL: sql, Graph: graph}, -1, nil
}
