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
package views

import (
	"log"

	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// ViewBean simple View model
type ViewBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Icon
	Icon string `json:"icon"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
	// URL
	URL string `json:"url"`
	// IsHome
	IsHome bool `json:"ishome"`
	// Devices
	Devices []core_models.IPersistent `json:"devices"`
}

// IViewBean interface
type IViewBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
}

// New constructor
func (p *ViewBean) New() IViewBean {
	bean := ViewBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *ViewBean) GetEntityName() string {
	return "ViewBean"
}

// Extend vars
func (p *ViewBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetID retrieve ID
func (p *ViewBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *ViewBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *ViewBean) Set(key string, value interface{}) {
	if key == "devices" {
		assert, ok := value.([]core_models.IPersistent)
		if ok {
			p.Devices = assert
		} else {
			log.Println("Warn: unable to assert type")
		}
	}
}

// SetString get set name
func (p *ViewBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *ViewBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *ViewBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// SetTimestamp set timestamp
func (p *ViewBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *ViewBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *ViewBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// ViewBeans simple bean model
type ViewBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []ViewBean
}

// New constructor
func (p *ViewBeans) New() core_models.IPersistents {
	bean := ViewBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]ViewBean, 0)}
	return &bean
}

// Add new bean
func (p *ViewBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, ViewBean{})
}

// Get collection of bean
func (p *ViewBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *ViewBeans) Index(index int) IViewBean {
	data, ok := p.Collection[index].(*ViewBean)
	if ok {
		return data
	}
	return nil
}
