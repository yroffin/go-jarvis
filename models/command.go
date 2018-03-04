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
	"encoding/json"
	"log"

	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// CommandHandler simple command handler
type CommandHandler interface {
	AsObject(*CommandBean, map[string]interface{}) (AsValue, error)
}

// CommandBean simple command model
type commandBean struct {
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
}

// CommandBean interface
type CommandBean interface {
	core_models.IPersistent
	// Value
	Set(string, interface{})
	SetString(string, string)
	GetAsString(string) string
	GetAsStringArray(string) []string
	ToString() string
	ToJSON() string
	// command
	GetType() string
}

// New constructor
func NewCommandBean() CommandBean {
	bean := commandBean{}
	return &bean
}

// SetName get set name
func (p *commandBean) SetName() string {
	return "Command"
}

// Type get set name
func (p *commandBean) GetType() string {
	return p.Type
}

// GetID retrieve ID
func (p *commandBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *commandBean) SetID(ID string) {
	p.ID = ID
}

// Set get set name
func (p *commandBean) Set(key string, value interface{}) {
}

// SetString get set name
func (p *commandBean) SetString(key string, value string) {
	switch key {
	case "body":
		p.Body = value
		break
	}
}

// Get get set name
func (p *commandBean) GetAsString(key string) string {
	switch key {
	case "body":
		return p.Body
	}
	return ""
}

// Get get set name
func (p *commandBean) GetAsStringArray(key string) []string {
	return make([]string, 0)
}

// ToString stringify this commnd
func (p *commandBean) ToString() string {
	payload, err := json.Marshal(p)
	if err != nil {
		log.Println("Unable to marshal:", err)
		return "{}"
	}
	return string(payload)
}

// ToJSON stringify this commnd
func (p *commandBean) ToJSON() string {
	payload, err := json.Marshal(p)
	if err != nil {
		log.Println("Unable to marshal:", err)
		return "{}"
	}
	return string(payload)
}

// SetTimestamp set timestamp
func (p *commandBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *commandBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *commandBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// CommandBeans simple bean model
type commandBeans struct {
	// Collection
	Collection []core_models.IPersistent
}

// NewCommandBeans constructor
func NewCommandBeans() core_models.IPersistents {
	bean := commandBeans{Collection: make([]core_models.IPersistent, 0)}
	return &bean
}

// Add new bean
func (p *commandBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
}

// Get collection of bean
func (p *commandBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *commandBeans) Index(index int) CommandBean {
	data, ok := p.Collection[index].(*commandBean)
	if ok {
		return data
	}
	return nil
}
