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
package process

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// ModelBean simple Model model
type ModelBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
}

// IModelBean interface
type IModelBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
}

// New constructor
func (p *ModelBean) New() IModelBean {
	bean := ModelBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *ModelBean) GetName() string {
	return "ModelBean"
}

// Extend vars
func (p *ModelBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetID retrieve ID
func (p *ModelBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *ModelBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *ModelBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *ModelBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *ModelBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *ModelBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// ToString stringify this commnd
func (p *ModelBean) ToString() string {
	// Call super method
	return core_models.IValueBean(p).ToString()
}

// ToJSON stringify this commnd
func (p *ModelBean) ToJSON() string {
	// Call super method
	return core_models.IValueBean(p).ToJSON()
}

// SetTimestamp set timestamp
func (p *ModelBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *ModelBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *ModelBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// ModelBeans simple bean model
type ModelBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []ModelBean
}

// New constructor
func (p *ModelBeans) New() core_models.IPersistents {
	bean := ModelBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]ModelBean, 0)}
	return &bean
}

// Add new bean
func (p *ModelBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, ModelBean{})
}

// Get collection of bean
func (p *ModelBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *ModelBeans) Index(index int) IModelBean {
	data, ok := p.Collection[index].(*ModelBean)
	if ok {
		return data
	}
	return nil
}
