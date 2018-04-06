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
package commands

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// CommandHandler simple command handler
type CommandHandler interface {
	AsObject(*CommandBean, map[string]interface{}) (core_models.ValueBean, error)
}

// CommandBean simple command model
type CommandBean struct {
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
	// Mode
	Mode string `json:"mode"`
	// Body
	Body string `json:"body"`
	// Extended internal store
	Extended map[string]interface{} `json:"extended"`
}

// ICommandBean interface
type ICommandBean interface {
	// inherit persistent behaviour
	core_models.IPersistent
	// inherit ValueBean behaviour
	core_models.IValueBean
	// command
	GetType() string
}

// New constructor
func (p *CommandBean) New() ICommandBean {
	bean := CommandBean{}
	bean.Extended = make(map[string]interface{})
	return &bean
}

// GetName get set name
func (p *CommandBean) GetName() string {
	return "CommandBean"
}

// Extend vars
func (p *CommandBean) Extend(e map[string]interface{}) {
	for k, v := range e {
		p.Extended[k] = v
	}
}

// GetType get set name
func (p *CommandBean) GetType() string {
	return p.Type
}

// GetID retrieve ID
func (p *CommandBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *CommandBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *CommandBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *CommandBean) SetString(key string, value string) {
	// Call super method
	core_models.IValueBean(p).SetString(key, value)
	switch key {
	case "body":
		p.Body = value
		break
	}
}

// Get get set name
func (p *CommandBean) GetAsString(key string) string {
	switch key {
	case "body":
		return p.Body
	default:
		// Call super method
		return core_models.IValueBean(p).GetAsString(key)
	}
}

// Get get set name
func (p *CommandBean) GetAsStringArray(key string) []string {
	// Call super method
	return core_models.IValueBean(p).GetAsStringArray(key)
}

// ToString stringify this commnd
func (p *CommandBean) ToString() string {
	// Call super method
	return core_models.IValueBean(p).ToString()
}

// ToJSON stringify this commnd
func (p *CommandBean) ToJSON() string {
	// Call super method
	return core_models.IValueBean(p).ToJSON()
}

// SetTimestamp set timestamp
func (p *CommandBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *CommandBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *CommandBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// CommandBeans simple bean model
type CommandBeans struct {
	// Collection
	Collection []core_models.IPersistent `json:"collections"`
	// Collection
	Collections []CommandBean
}

// New constructor
func (p *CommandBeans) New() core_models.IPersistents {
	bean := CommandBeans{Collection: make([]core_models.IPersistent, 0), Collections: make([]CommandBean, 0)}
	return &bean
}

// Add new bean
func (p *CommandBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
	p.Collections = append(p.Collections, CommandBean{})
}

// Get collection of bean
func (p *CommandBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *CommandBeans) Index(index int) ICommandBean {
	data, ok := p.Collection[index].(*CommandBean)
	if ok {
		return data
	}
	return nil
}