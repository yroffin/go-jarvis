// Package bpmn module
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
package bpmn

import (
	"crypto/rand"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-jarvis/core/apis"
)

// Engine internal members
type Engine struct {
	// internal members
	Name string
	// Process
	specifications map[string]*ProcessBpmn
	// Process
	subcribers []func(IBpmnEngine, *Event) error
	// Registry
	registry map[string]*ProcessInstance
}

// IBpmnEngine implements IBean
type IBpmnEngine interface {
	Deploy(bpmn string) (*ProcessBpmn, error)
	CreateWorkflow(id string, maylod map[string]interface{}) (*ProcessInstance, error)
	Subscribe(func(IBpmnEngine, *Event) error) error
	Start(*ProcessInstance) (*ProcessInstance, error)
	Active() apis.Graph
}

// New constructor
func (p *Engine) New() IBpmnEngine {
	bean := Engine{}
	bean.subcribers = make([]func(IBpmnEngine, *Event) error, 0)
	bean.specifications = make(map[string]*ProcessBpmn)
	bean.registry = make(map[string]*ProcessInstance)
	go bean.Heartbeat()
	return &bean
}

// UUID generates a random UUID according to RFC 4122
func UUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "N/A"
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	var text = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return text
}

// Deploy API
func (p *Engine) Deploy(bpmn string) (*ProcessBpmn, error) {
	// Read bpmn structure
	decode := bpmnXML{}
	xml.Unmarshal([]byte(bpmn), &decode)
	log.WithFields(log.Fields{
		"Body": bpmn,
		"Xml":  decode,
	}).Debug("Deploy")
	// Process
	proc := ProcessBpmn{
		ID:     decode.Process.ID,
		Tasks:  make(map[string]TaskBpmn),
		Events: make(map[string]EventBpmn),
		Flows:  make(map[string]FlowBpmn),
	}
	// find all service
	var services = make([]string, 0)
	for _, start := range decode.Process.Start {
		for _, out := range start.Out {
			// Store events
			event := EventBpmn{
				ID:  start.ID,
				Out: out.Body,
			}
			proc.Events[start.ID] = event
		}
	}
	for _, end := range decode.Process.End {
		for _, inc := range end.Inc {
			// Store events
			event := EventBpmn{
				ID:  end.ID,
				Inc: inc.Body,
			}
			proc.Events[end.ID] = event
		}
	}
	for _, flow := range decode.Process.Flow {
		// Store flows
		flow := FlowBpmn{
			ID:     flow.ID,
			Source: flow.Source,
			Target: flow.Target,
		}
		proc.Flows[flow.ID] = flow
	}
	for _, svc := range decode.Process.Service {
		// Store tasks
		task := TaskBpmn{
			ID: svc.ID,
		}
		proc.Tasks[svc.ID] = task
		for _, ext := range svc.Extention {
			for _, svcType := range ext.TaskDefinition {
				services = append(services, svcType.Type)
			}
		}
	}
	// Store specification
	p.specifications[proc.ID] = &proc
	return &proc, nil
}

// CreateWorkflow API
func (p *Engine) CreateWorkflow(id string, payload map[string]interface{}) (*ProcessInstance, error) {
	instance := ProcessInstance{}
	instance.ID = UUID()
	for _, find := range p.specifications {
		if find.ID == id {
			b, _ := json.Marshal(find)
			copy := ProcessBpmn{}
			json.Unmarshal(b, &copy)
			instance.Definition = copy
			instance.Nodes = []NodeInstance{}
			instance.Edges = []EdgeInstance{}
		}
	}
	p.registry[instance.ID] = &instance
	return &instance, nil
}

// Subscribe API
func (p *Engine) Subscribe(handler func(p IBpmnEngine, event *Event) error) error {
	p.subcribers = append(p.subcribers, handler)
	return nil
}

// Start the instance
func (p *Engine) Start(proc *ProcessInstance) (*ProcessInstance, error) {
	p.registry[proc.ID].Start()
	return proc, nil
}

// Heartbeat daemon proc
func (p *Engine) Heartbeat() {
	var counter = 0
	for true {
		var active = 0
		var inactive = 0
		var list = []string{}
		// Scan current status
		for key, instance := range p.registry {
			if instance.Active {
				active++
				list = append(list, key)
			} else {
				inactive++
			}
		}
		counter++
		// Run active
		for _, instance := range p.registry {
			if instance.Active {
				for _, node := range p.Run(instance) {
					event := Event{
						Process: instance,
						Node:    node,
					}
					for _, subscribed := range p.subcribers {
						subscribed(p, &event)
						instance.CompleteByID(event.Node.ID, true)
					}
				}
			}
		}
		time.Sleep(3000 * time.Millisecond)
	}
}

// Run daemon proc
func (p *Engine) Run(instance *ProcessInstance) []*NodeInstance {
	return instance.Run()
}

// Active get all active instance
func (p *Engine) Active() apis.Graph {
	var options = []byte(`
		{
			"configure": {
				"enabled": true
			},
			"nodes": {
			  "color": {
				"highlight": {},
				"hover": {}
			  },
			  "font": {
			  },
			  "scaling": {
			  },
			  "shapeProperties": {
			  }
			},
			"edges": {
			  "smooth": false
			},
			"layout": {
			  "hierarchical": {
				"enabled": true,
				"levelSeparation": -150,
				"direction": "DU"
			  }
			},
			"interaction": {
			  "hover": true
			},
			"physics": {
			  "hierarchicalRepulsion": {
				"centralGravity": 0
			  },
			  "minVelocity": 0.75,
			  "solver": "hierarchicalRepulsion"
			},
			"groups": {
				"Event": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf1db",
						"size": 32,
						"color": "#000000"
					}
				},
				"Task": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf0ae",
						"size": 32,
						"color": "#000000"
					}
				},
				"Event::Active": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf1db",
						"size": 64,
						"color": "#FF0000"
					}
				},
				"Task::Active": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf0ae",
						"size": 64,
						"color": "#FF0000"
					}
				}
			}
		  }`)
	graph := apis.Graph{
		Nodes:   make([]apis.Node, 0),
		Edges:   make([]apis.Edge, 0),
		Options: map[string]interface{}{},
	}
	err := json.Unmarshal(options, &graph.Options)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Options")
	}
	// Load graph
	for _, instance := range p.registry {
		for _, node := range instance.Nodes {
			var group = ""
			if node.IsEvent {
				group = "Event"
				if node.ID == instance.Current {
					group = "Event::Active"
				}
			} else {
				group = "Task"
				if node.ID == instance.Current {
					group = "Task::Active"
				}
			}
			graph.Nodes = append(graph.Nodes, apis.Node{
				ID:    node.ID,
				Label: node.Ref,
				Group: group,
				Title: node.ID,
			})
		}
		for _, edge := range instance.Edges {
			graph.Edges = append(graph.Edges, apis.Edge{
				Label: edge.ID,
				From:  edge.Source,
				To:    edge.Target,
				Data:  "{}",
			})
		}
	}
	return graph
}
