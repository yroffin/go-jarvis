// Package models for all models
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
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// ScriptPluginBean simple ScriptPlugin model
type ScriptPluginBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Icon
	Icon string `json:"icon"`
	// Type
	Type string `json:"type"`
	// Owner
	Owner string `json:"owner"`
	// Active
	Active bool `json:"active"`
	// Visible
	Visible bool `json:"visible"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
}

// IScriptPluginBean interface
type IScriptPluginBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
}

// New constructor
func (p *ScriptPluginBean) New() IScriptPluginBean {
	bean := ScriptPluginBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *ScriptPluginBean) GetName() string {
	return "ScriptPluginBean"
}

// Extend vars
func (p *ScriptPluginBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetID retrieve ID
func (p *ScriptPluginBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *ScriptPluginBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *ScriptPluginBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *ScriptPluginBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *ScriptPluginBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *ScriptPluginBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// ToString stringify this commnd
func (p *ScriptPluginBean) ToString() string {
	// Call super method
	return core_models.IValueBean(p).ToString()
}

// ToJSON stringify this commnd
func (p *ScriptPluginBean) ToJSON() string {
	// Call super method
	return core_models.IValueBean(p).ToJSON()
}

// SetTimestamp set timestamp
func (p *ScriptPluginBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *ScriptPluginBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *ScriptPluginBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// ScriptPluginBeans simple bean model
type ScriptPluginBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []ScriptPluginBean
}

// New constructor
func (p *ScriptPluginBeans) New() core_models.IPersistents {
	bean := ScriptPluginBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]ScriptPluginBean, 0)}
	return &bean
}

// Add new bean
func (p *ScriptPluginBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, ScriptPluginBean{})
}

// Get collection of bean
func (p *ScriptPluginBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *ScriptPluginBeans) Index(index int) IScriptPluginBean {
	data, ok := p.Collection[index].(*ScriptPluginBean)
	if ok {
		return data
	}
	return nil
}
