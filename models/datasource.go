// Package DataSources for all DataSources
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
package models

import core_models "github.com/yroffin/go-boot-sqllite/core/models"

// DataSourceBean simple DataSource DataSource
type DataSourceBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Icon
	Icon string `json:"icon"`
	// Collect
	Collect string `json:"collect"`
	// Pipes
	Pipes string `json:"pipes"`
	// Body
	Body string `json:"body"`
	// Extended
	Extended map[string]interface{} `json:"extended"`
}

// IDataSourceBean interface
type IDataSourceBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
}

// New constructor
func (p *DataSourceBean) New() IDataSourceBean {
	bean := DataSourceBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *DataSourceBean) GetName() string {
	return "DataSourceBean"
}

// Extend vars
func (p *DataSourceBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetID retrieve ID
func (p *DataSourceBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *DataSourceBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *DataSourceBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *DataSourceBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *DataSourceBean) GetAsString(key string) string {
	switch key {
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *DataSourceBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// ToString stringify this commnd
func (p *DataSourceBean) ToString() string {
	// Call super method
	return core_models.IValueBean(p).ToString()
}

// ToJSON stringify this commnd
func (p *DataSourceBean) ToJSON() string {
	// Call super method
	return core_models.IValueBean(p).ToJSON()
}

// SetTimestamp set timestamp
func (p *DataSourceBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *DataSourceBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *DataSourceBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// DataSourceBeans simple bean DataSource
type DataSourceBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []DataSourceBean
}

// New constructor
func (p *DataSourceBeans) New() core_models.IPersistents {
	bean := DataSourceBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]DataSourceBean, 0)}
	return &bean
}

// Add new bean
func (p *DataSourceBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, DataSourceBean{})
}

// Get collection of bean
func (p *DataSourceBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *DataSourceBeans) Index(index int) IDataSourceBean {
	data, ok := p.Collection[index].(*DataSourceBean)
	if ok {
		return data
	}
	return nil
}
