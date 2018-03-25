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
package models

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// NotificationBean simple model
type NotificationBean struct {
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
}

// INotificationBean interface
type INotificationBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
	// command
	GetType() string
}

// New constructor
func (p *NotificationBean) New() INotificationBean {
	bean := NotificationBean{}
	return &bean
}

// SetName get set name
func (p *NotificationBean) SetName() string {
	return "Notification"
}

// GetType get set name
func (p *NotificationBean) GetType() string {
	return p.Type
}

// GetID retrieve ID
func (p *NotificationBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *NotificationBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *NotificationBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *NotificationBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
}

// Get get set name
func (p *NotificationBean) GetAsString(key string) string {
	// Call super method
	return core_models.IValueBean(p).GetAsString(key)
}

// Get get set name
func (p *NotificationBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// ToString stringify this commnd
func (p *NotificationBean) ToString() string {
	// Call super method
	return core_models.IValueBean(p).ToString()
}

// ToJSON stringify this commnd
func (p *NotificationBean) ToJSON() string {
	// Call super method
	return core_models.IValueBean(p).ToJSON()
}

// SetTimestamp set timestamp
func (p *NotificationBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *NotificationBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *NotificationBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// NotificationBeans simple bean model
type NotificationBeans struct {
	// Collection
	Collection []core_models.IPersistent
}

// New constructor
func (p *NotificationBeans) New() core_models.IPersistents {
	bean := NotificationBeans{Collection: make([]core_models.IPersistent, 0)}
	return &bean
}

// Add new bean
func (p *NotificationBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
}

// Get collection of bean
func (p *NotificationBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *NotificationBeans) Index(index int) *NotificationBean {
	data, ok := p.Collection[index].(*NotificationBean)
	if ok {
		return data
	}
	return nil
}
