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
package devices

import (
	"encoding/json"

	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// DeviceBean simple Device model
type DeviceBean struct {
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
	// Parameters
	Parameters string `json:"parameters"`
	// Owner
	Owner string `json:"owner"`
	// Visible
	Visible bool `json:"visible"`
	// TagColor
	TagColor string `json:"tagColor"`
	// TagOpacity
	TagOpacity string `json:"tagOpacity"`
	// TagTextColor
	TagTextColor string `json:"tagTextColor"`
	// RowSpan
	RowSpan string `json:"rowSpan"`
	// ColSpan
	ColSpan string `json:"colSpan"`
	// Template
	Template string `json:"template"`
}

// IDeviceBean interface
type IDeviceBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
	// GetParameters
	GetParameters() map[string]interface{}
}

// New constructor
func (p *DeviceBean) New() IDeviceBean {
	bean := DeviceBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetEntityName get set name
func (p *DeviceBean) GetEntityName() string {
	return "DeviceBean"
}

// Extend vars
func (p *DeviceBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetExtend vars
func (p *DeviceBean) GetExtend() map[string]interface{} {
	return p.Extended
}

// GetID retrieve ID
func (p *DeviceBean) GetID() string {
	return p.ID
}

// GetParameters field
func (p *DeviceBean) GetParameters() map[string]interface{} {
	value := make(map[string]interface{})
	json.Unmarshal([]byte(p.Parameters), &value)
	return value
}

// SetID retrieve ID
func (p *DeviceBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *DeviceBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *DeviceBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// GetAsString method
func (p *DeviceBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// GetAsStringArray method
func (p *DeviceBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// SetTimestamp set timestamp
func (p *DeviceBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *DeviceBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *DeviceBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// DeviceBeans simple bean model
type DeviceBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []DeviceBean
}

// New constructor
func (p *DeviceBeans) New() core_models.IPersistents {
	bean := DeviceBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]DeviceBean, 0)}
	return &bean
}

// Add new bean
func (p *DeviceBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, DeviceBean{})
}

// Get collection of bean
func (p *DeviceBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *DeviceBeans) Index(index int) IDeviceBean {
	data, ok := p.Collection[index].(*DeviceBean)
	if ok {
		return data
	}
	return nil
}
