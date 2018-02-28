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

// SetName get set name
func (p *NotificationBean) SetName() string {
	return "Notification"
}

// GetID retrieve ID
func (p *NotificationBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *NotificationBean) SetID(ID string) {
	p.ID = ID
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
