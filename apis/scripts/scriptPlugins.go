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
package scripts

import (
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/apis/commands"
)

func init() {
	winter.Helper.Register("ScriptPluginBean", (&ScriptPlugin{}).New())
}

// ScriptPlugin internal members
type ScriptPlugin struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/plugins/scripts"`
	// LinkCommand with injection mecanism
	LinkCommand commands.ICommand `@autowired:"CommandBean" @link:"/api/plugins/scripts" @href:"commands"`
	Command     commands.ICommand `@autowired:"CommandBean"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IScriptPlugin implements IBean
type IScriptPlugin interface {
	engine.IAPI
}

// New constructor
func (p *ScriptPlugin) New() IScriptPlugin {
	bean := ScriptPlugin{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *ScriptPlugin) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ScriptPluginBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ScriptPluginBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *ScriptPlugin) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *ScriptPlugin) Validate(name string) error {
	return nil
}
