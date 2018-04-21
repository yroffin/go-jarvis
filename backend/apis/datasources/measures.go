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
package datasources

import (
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/apis/connectors"
)

func init() {
	winter.Helper.Register("MeasureBean", (&Measure{}).New())
}

// Measure internal members
type Measure struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/measures"`
	// LinkConnector with injection mecanism
	LinkConnector connectors.IConnector `@autowired:"ConnectorBean" @link:"/api/measures" @href:"connectors"`
	Connector     connectors.IConnector `@autowired:"ConnectorBean"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IMeasure implements IBean
type IMeasure interface {
	engine.IAPI
}

// New constructor
func (p *Measure) New() IMeasure {
	bean := Measure{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Measure) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&MeasureBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&MeasureBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Measure) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Measure) Validate(name string) error {
	return nil
}
