// Package interfaces for common interfaces
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

	log "github.com/sirupsen/logrus"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services/cron"
)

func init() {
	winter.Helper.Register("CronBean", (&Cron{}).New())
}

// Cron internal members
type Cron struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/crons"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// Cron with injection mecanism
	Cron cron.ICronService `@autowired:"cron-service"`
	// Cron with injection mecanism
	Trigger ITrigger `@autowired:"TriggerBean"`
}

// ICron implements IBean
type ICron interface {
	engine.IAPI
}

// New constructor
func (p *Cron) New() ICron {
	bean := Cron{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Cron) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&CronBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&CronBeans{}).New()
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		var parameters = make(map[string]interface{})
		json.Unmarshal([]byte(body), &parameters)
		if name == "toggle" {
			// task
			return p.Toggle(id, parameters)
		}
		if name == "test" {
			// task
			return p.Test(id, parameters)
		}
		return parameters, -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Cron) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Cron) Validate(name string) error {
	return nil
}

// Toggle this cron
func (p *Cron) Toggle(id string, parameters map[string]interface{}) (interface{}, int, error) {
	// Search in current cron
	exist := p.Cron.Exist(id)
	// Retrieve parameters
	if exist {
		err := p.Cron.Toggle(id)
		log.WithFields(log.Fields{
			"job": id,
			"err": err,
		}).Info("Job does exist, togglt it")
	} else {
		// get crontab
		crontab, _ := p.GetByID(id)
		log.WithFields(log.Fields{
			"job":         id,
			"triggerType": crontab.(ICronBean).GetTriggerType(),
		}).Warn("Job does not exist")
		switch crontab.(ICronBean).GetTriggerType() {
		case "CRONTAB":
			if len(crontab.(ICronBean).GetCron()) <= 0 {
				return false, -1, nil
			}
			p.Cron.Add(id, crontab.(ICronBean).GetCron(), func() {
				// Retrieve event
				events := p.eventFactory(id)
				for _, event := range events {
					// Post a trigger
					p.Trigger.Post(event)
				}
			})
			break
		default:
			break
		}
	}
	return exist, -1, nil
}

// Test this cron
func (p *Cron) Test(id string, parameters map[string]interface{}) (interface{}, int, error) {
	// Retrieve event
	events := p.eventFactory(id)
	for _, event := range events {
		// Post a trigger
		p.Trigger.Post(event)
	}
	return events, len(events), nil
}

// Build event to send
func (p *Cron) eventFactory(id string) []EventBean {
	// Triggers
	result := make([]EventBean, 0)
	// Retrieve parameters
	crontab, _ := p.GetByID(id)
	// Retrieve all triggers
	triggers, _ := p.Trigger.GetAll()
	for _, trigger := range triggers {
		linked, _ := p.Trigger.GetAllLinks(trigger.GetID(), winter.Helper.GetBean("CronBean").(engine.IAPI))
		for _, cron := range linked {
			if cron.GetID() == crontab.GetID() {
				// Post a trigger
				result = append(result, EventNew(trigger.GetID(), crontab.(ICronBean).GetName()))
			}
		}
	}
	return result
}
