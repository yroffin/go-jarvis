// Package mqtt for common interfaces
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
package cron

import (
	crontab "github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("cron-service", (&service{}).New())
}

// service internal members
type service struct {
	// members
	*winter.Service
	// Service map
	services map[string]*job
	// Service cron
	crontab *crontab.Cron
}

// internal job
type job struct {
	// JobID
	id string
	// Active
	active bool
	// Counter
	counter int
	// Param
	param string
	// Handler
	handler func()
}

// ICronService implements IBean
type ICronService interface {
	winter.IService
	Exist(string) bool
	Toggle(key string) error
	Add(key string, param string, handler func()) error
}

// New constructor
func (p *service) New() ICronService {
	bean := service{Service: &winter.Service{Bean: &winter.Bean{}}}
	bean.services = make(map[string]*job)
	bean.crontab = crontab.New()
	bean.crontab.Start()
	return &bean
}

// Init this API
func (p *service) Init() error {
	return nil
}

// PostConstruct this API
func (p *service) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *service) Validate(name string) error {
	return nil
}

// Check for existing job
func (p *service) Exist(key string) bool {
	_, find := p.services[key]
	if find {
		return true
	} else {
		return false
	}
}

// Run this job
func (p *job) Run() {
	p.counter = p.counter + 1
	if p.active {
		p.handler()
	} else {
		// Warn a little
		if p.counter < 2 {
			log.WithFields(log.Fields{
				"job": p.id,
			}).Warn("Job is not active")
		}

	}
}

// toggle this job
func (p *job) toggle(c *service) {
	p.active = !p.active
	p.counter = 0
}

// Check for existing job
func (p *service) Add(key string, param string, handler func()) error {
	if !p.Exist(key) {
		job := job{id: key, param: param, active: true, handler: handler}
		p.services[key] = &job
		// Add this job (with restart in order to reschedule)
		p.crontab.Stop()
		p.crontab.AddJob(param, &job)
		p.crontab.Start()
		log.WithFields(log.Fields{
			"job":   key,
			"param": param,
		}).Info("Submitted")
	} else {
		log.WithFields(log.Fields{
			"job": key,
		}).Warn("Already submitted")
	}
	return nil
}

// Check for existing job
func (p *service) Toggle(key string) error {
	if p.Exist(key) {
		p.services[key].toggle(p)
	}
	return nil
}
