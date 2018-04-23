// Package events for common interfaces
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
	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("TriggerBean", (&Trigger{}).New())
}

// Trigger internal members
type Trigger struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/triggers"`
	// LinkCron with injection mecanism
	LinkCron ICron `@autowired:"CronBean" @link:"/api/triggers" @href:"crons"`
	// Device with injection mecanism
	Device IDevice `@autowired:"DeviceBean"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// Internal channel for events
	events chan EventBean
}

// ITrigger implements IBean
type ITrigger interface {
	engine.IAPI
	// Post a new event
	Post(event EventBean) error
}

// New constructor
func (p *Trigger) New() ITrigger {
	bean := Trigger{API: &engine.API{Bean: &winter.Bean{}}}
	bean.events = make(chan EventBean)
	go bean.Handler()
	bean.GetByIDListener = make([]func(models.IPersistent) models.IPersistent, 1)
	bean.GetByIDListener[0] = bean.middleware
	return &bean
}

// Init this API
func (p *Trigger) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&TriggerBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&TriggerBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Trigger) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Trigger) Validate(name string) error {
	return nil
}

// Post a new event
func (p *Trigger) Post(event EventBean) error {
	p.events <- event
	return nil
}

// Simple middle ware to override status
func (p *Trigger) middleware(entity models.IPersistent) models.IPersistent {
	result := make([]models.IPersistent, 0)
	// Iterate on devices, then triggers and find attached device
	devices, _ := p.Device.GetAll()
	for _, device := range devices {
		linked, _ := p.Device.(IDevice).GetAllLinks(device.GetID(), winter.Helper.GetBean("TriggerBean").(engine.IAPI))
		for _, trigger := range linked {
			if entity.GetID() == trigger.GetID() {
				result = append(result, device)
			}
		}
	}
	// Update device
	entity.(*TriggerBean).Devices = result
	return entity
}

// Handler events
func (p *Trigger) Handler() error {
	for {
		event := <-p.events
		log.WithFields(log.Fields{
			"id":   event.ID,
			"text": event.Text,
		}).Info("Event handler")

		// Apply event
		for _, device := range p.devices(event) {
			log.WithFields(log.Fields{
				"id":   device.GetID(),
				"name": device.GetName(),
			}).Info("Event handler")
		}
	}
}

// Handler events on device
func (p *Trigger) devices(event EventBean) []IDeviceBean {
	filtered := make(map[string]IDeviceBean)
	devices, _ := p.Device.GetAll()
	for _, device := range devices {
		triggers, _ := p.Device.GetAllLinks(device.GetID(), winter.Helper.GetBean("TriggerBean").(engine.IAPI))
		for _, trigger := range triggers {
			if trigger.GetID() == event.ID {
				filtered[device.GetID()] = device.(IDeviceBean)
			}
		}
	}
	// Build uniq
	uniq := make([]IDeviceBean, 0)
	for _, device := range filtered {
		uniq = append(uniq, device)
	}
	return uniq
}
