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
package system

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// SnapshotBean simple Snapshot model
type SnapshotBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Type
	Type string `json:"type"`
	// Icon
	Icon string `json:"icon"`
	// Json
	JSON interface{} `json:"json"`
	// Extended internal store
	Extended map[string]interface{} `json:"extended"`
}

// ISnapshotBean interface
type ISnapshotBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
	// Snapshot
	GetType() string
	// GetJSON
	GetJSON() interface{}
}

// New constructor
func (p *SnapshotBean) New() ISnapshotBean {
	bean := SnapshotBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *SnapshotBean) GetName() string {
	return "SnapshotBean"
}

// Extend vars
func (p *SnapshotBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetType get set name
func (p *SnapshotBean) GetType() string {
	return p.Type
}

// GetJSON get set json
func (p *SnapshotBean) GetJSON() interface{} {
	return p.JSON
}

// GetID retrieve ID
func (p *SnapshotBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *SnapshotBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *SnapshotBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *SnapshotBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
	switch key {
	}
}

// GetAsString get as string
func (p *SnapshotBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// GetAsStringArray get array
func (p *SnapshotBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// SetTimestamp set timestamp
func (p *SnapshotBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *SnapshotBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *SnapshotBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// SnapshotBeans simple bean model
type SnapshotBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []SnapshotBean
}

// New constructor
func (p *SnapshotBeans) New() core_models.IPersistents {
	bean := SnapshotBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]SnapshotBean, 0)}
	return &bean
}

// Add new bean
func (p *SnapshotBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, SnapshotBean{})
}

// Get collection of bean
func (p *SnapshotBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *SnapshotBeans) Index(index int) ISnapshotBean {
	data, ok := p.Collection[index].(*SnapshotBean)
	if ok {
		return data
	}
	return nil
}
