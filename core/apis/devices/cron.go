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
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// CronBean simple Cron model
type CronBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Icon
	Icon string `json:"icon"`
	// TriggerType
	TriggerType string `json:"triggerType"`
	// Latitude
	Latitude string `json:"latitude"`
	// Longitude
	Longitude string `json:"longitude"`
	// Shift
	Shift int64 `json:"shift"`
	// Cron
	Cron string `json:"cron"`
	// StartAtRuntime
	StartAtRuntime bool `json:"startAtRuntime"`
	// Status
	Status bool `json:"status"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
}

// ICronBean interface
type ICronBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
	// GetTriggerType cron type
	GetTriggerType() string
	// Get cron string
	GetCron() string
	// Name
	GetName() string
}

// New constructor
func (p *CronBean) New() ICronBean {
	bean := CronBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetEntityName get set name
func (p *CronBean) GetEntityName() string {
	return "CronBean"
}

// GetName get name
func (p *CronBean) GetName() string {
	return p.Name
}

// Extend vars
func (p *CronBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetExtend vars
func (p *CronBean) GetExtend() map[string]interface{} {
	return p.Extended
}

// GetTriggerType retrieve cron string
func (p *CronBean) GetTriggerType() string {
	return p.TriggerType
}

// GetCron retrieve cron string
func (p *CronBean) GetCron() string {
	return p.Cron
}

// GetID retrieve ID
func (p *CronBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *CronBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *CronBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *CronBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *CronBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *CronBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// SetTimestamp set timestamp
func (p *CronBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *CronBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *CronBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// CronBeans simple bean model
type CronBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []CronBean
}

// New constructor
func (p *CronBeans) New() core_models.IPersistents {
	bean := CronBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]CronBean, 0)}
	return &bean
}

// Add new bean
func (p *CronBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, CronBean{})
}

// Get collection of bean
func (p *CronBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *CronBeans) Index(index int) ICronBean {
	data, ok := p.Collection[index].(*CronBean)
	if ok {
		return data
	}
	return nil
}
