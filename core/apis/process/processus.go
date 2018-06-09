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
package process

import (
	"encoding/json"

	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/apis/devices"
	"github.com/yroffin/go-jarvis/core/bpmn"

	"github.com/zeebe-io/zbc-go/zbc"
)

func init() {
	winter.Helper.Register("ProcessBean", (&Processus{}).New())
}

// Processus internal members
type Processus struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/processes"`
	// Trigger with injection mecanism
	LinkTrigger devices.ITrigger `@autowired:"TriggerBean" @link:"/api/processes" @href:"triggers"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// client zeebe
	zbClient *zbc.Client
	// All processus
	engine bpmn.IBpmnEngine
}

// IProcessus implements IBean
type IProcessus interface {
	engine.IAPI
}

// New constructor
func (p *Processus) New() IProcessus {
	bean := Processus{API: &engine.API{Bean: &winter.Bean{}}}
	bean.engine = (&bpmn.Engine{}).New()
	return &bean
}

// Init this API
func (p *Processus) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ProcessusBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ProcessusBeans{}).New()
	}
	p.HandlerTasks = func(name string, body string) (interface{}, int, error) {
		if name == "active" {
			// task
			return p.Active(body)
		}
		return "", -1, nil
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		var parameters = make(map[string]interface{})
		json.Unmarshal([]byte(body), &parameters)
		if name == "deploy" {
			// task
			return p.Deploy(id, []byte(body))
		}
		if name == "execute" {
			// task
			return p.Execute(id, []byte(body))
		}
		return parameters, -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Processus) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Processus) Validate(name string) error {
	// publish subscriber
	p.engine.Subscribe(handler)
	// deploy all current process
	var all, _ = p.GetAll()
	for _, proc := range all {
		p.Deploy(proc.GetID(), []byte{})
	}
	return nil
}

func handler(engine bpmn.IBpmnEngine, event *bpmn.Event) error {
	fmt.Printf("Event: %v\n", models.ToJSON(event.Node))
	event.Node.IsCompleted = true
	return nil
}

// Deploy API
func (p *Processus) Deploy(id string, body []byte) (interface{}, int, error) {
	// Retrieve entity
	entity, _ := p.GetByID(id)
	// Deploy proc
	response, err := p.engine.Deploy(entity.(IProcessusBean).GetBpmn())
	// Check for error
	if err != nil {
		log.WithFields(log.Fields{
			"Detail": err,
			"Xml":    entity.(IProcessusBean).GetBpmn(),
		}).Warn("Deploy")
	} else {
		log.WithFields(log.Fields{
			"Process": response,
		}).Debug("Deploy")
		// Fix process name based on xml definition
		entity.(*ProcessusBean).Name = response.ID
		p.SQLCrudBusiness.Update(entity)
	}
	return response, -1, nil
}

// Execute API
func (p *Processus) Execute(id string, body []byte) (interface{}, int, error) {
	// After the workflow is deployed.
	payload := make(map[string]interface{})
	json.Unmarshal(body, &payload)

	// Retrieve entity
	entity, _ := p.GetByID(id)
	response, err := p.engine.CreateWorkflow(entity.(*ProcessusBean).Name, payload)
	if err != nil {
		log.WithFields(log.Fields{
			"Detail": err,
			"Name":   entity.(*ProcessusBean).Name,
		}).Warn("execute")
	}

	p.engine.Start(response)

	return response, -1, nil
}

// Active get all active instance as graph
func (p *Processus) Active(body string) (interface{}, int, error) {
	return p.engine.Active(), -1, nil
}
