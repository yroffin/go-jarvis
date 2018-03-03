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
}

// SetName get set name
func (p *CommandBean) SetName() string {
	return "Command"
}

// GetID retrieve ID
func (p *CommandBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *CommandBean) SetID(ID string) {
	p.ID = ID
}

// ToString stringify this commnd
func (p *CommandBean) ToString() string {
	payload, err := json.Marshal(p)
	if err != nil {
		log.Println("Unable to marshal:", err)
		return "{}"
	}
	return string(payload)
}

// ToJSON stringify this commnd
func (p *CommandBean) ToJSON() string {
	payload, err := json.Marshal(p)
	if err != nil {
		log.Println("Unable to marshal:", err)
		return "{}"
	}
	return string(payload)
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
	Collection []core_models.IPersistent
}

// Add new bean
func (p *CommandBeans) Add(bean core_models.IPersistent) {
	p.Collection = append(p.Collection, bean)
}

// Get collection of bean
func (p *CommandBeans) Get() []core_models.IPersistent {
	return p.Collection
}

// Index read a single element
func (p *CommandBeans) Index(index int) *CommandBean {
	data, ok := p.Collection[index].(*CommandBean)
	if ok {
		return data
	}
	return nil
}
