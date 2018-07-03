// Package apis for common apis
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
package system

import (
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/auto"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
)

var (
	version = "1.0"
)

func init() {
	winter.Helper.Register("SecurityBean", (&Security{}).New())
}

// Security internal members
type Security struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// Api
	SecConnect interface{} `@handler:"Connect" path:"/api/connect" method:"GET" mime-type:"/application/json"`
	SecProfile interface{} `@handler:"Profile" path:"/api/profile/me" method:"GET" mime-type:"/application/json"`
	SecVersion interface{} `@handler:"Version" path:"/api/version" method:"GET" mime-type:"/application/json"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// ISecurity implements IBean
type ISecurity interface {
	engine.IAPI
}

// New constructor
func (p *Security) New() ISecurity {
	bean := Security{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Security) Init() error {
	return p.API.Init()
}

// PostConstruct this API
func (p *Security) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Security) Validate(name string) error {
	return nil
}

// Connect API
func (p *Security) Connect() func() (string, error) {
	anonymous := func() (string, error) {
		return "false", nil
	}
	return anonymous
}

// Profile API
func (p *Security) Profile() func() (string, error) {
	anonymous := func() (string, error) {
		return "{\"attributes\":{\"email\":\"-\"}}", nil
	}
	return anonymous
}

// Version API
func (p *Security) Version() func() (string, error) {
	anonymous := func() (string, error) {
		pack := auto.Pack
		value, _ := pack.MustString("version/version.json")
		return value, nil
	}
	return anonymous
}
