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
package events

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// TriggerBean simple Trigger model
type TriggerBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Icon
	Icon string `json:"icon"`
	// Topic
	Topic string `json:"topic"`
	// Body
	Body string `json:"body"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
}

// ITriggerBean interface
type ITriggerBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
}

// New constructor
func (p *TriggerBean) New() ITriggerBean {
	bean := TriggerBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *TriggerBean) GetEntityName() string {
	return "TriggerBean"
}

// Extend vars
func (p *TriggerBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetExtend vars
func (p *TriggerBean) GetExtend() map[string]interface{} {
	return p.Extended
}

// GetID retrieve ID
func (p *TriggerBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *TriggerBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *TriggerBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *TriggerBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *TriggerBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *TriggerBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// SetTimestamp set timestamp
func (p *TriggerBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *TriggerBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *TriggerBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// TriggerBeans simple bean model
type TriggerBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []TriggerBean
}

// New constructor
func (p *TriggerBeans) New() core_models.IPersistents {
	bean := TriggerBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]TriggerBean, 0)}
	return &bean
}

// Add new bean
func (p *TriggerBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, TriggerBean{})
}

// Get collection of bean
func (p *TriggerBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *TriggerBeans) Index(index int) ITriggerBean {
	data, ok := p.Collection[index].(*TriggerBean)
	if ok {
		return data
	}
	return nil
}
