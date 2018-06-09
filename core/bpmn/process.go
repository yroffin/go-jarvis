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

import log "github.com/sirupsen/logrus"

// Start the instance
func (p *ProcessInstance) Start() {
	if !p.Compiled {
		// Find start event, and others
		for _, event := range p.Definition.Events {
			if len(event.Out) > 0 {
				p.Current = UUID()
				p.Nodes = append(p.Nodes, NodeInstance{
					ID:          p.Current,
					Ref:         event.ID,
					IsEvent:     true,
					IsStart:     true,
					IsCompleted: true,
				})
			}
			if len(event.Inc) > 0 {
				p.Nodes = append(p.Nodes, NodeInstance{
					ID:      UUID(),
					Ref:     event.ID,
					IsEvent: true,
				})
			}
		}
		// Find all tasks
		for _, task := range p.Definition.Tasks {
			p.Nodes = append(p.Nodes, NodeInstance{
				ID:     UUID(),
				Ref:    task.ID,
				IsTask: true,
			})
		}
		// Find all flows
		for _, flow := range p.Definition.Flows {
			var source, target string
			for _, node := range p.Nodes {
				if node.Ref == flow.Source {
					source = node.ID
				}
				if node.Ref == flow.Target {
					target = node.ID
				}
			}
			p.Edges = append(p.Edges, EdgeInstance{
				ID:     UUID(),
				Source: source,
				Target: target,
			})
		}
		p.Compiled = true
	}
	p.Active = true
}

// Run the instance
func (p *ProcessInstance) Run() []*NodeInstance {
	_, target := p.FindNodeByID(p.Current)
	if target.IsCompleted {
		hasNext, next := p.NextNode(p.Current)
		if hasNext {
			p.Current = next.ID
			for _, node := range p.Nodes {
				if node.ID == p.Current {
					node.IsActive = true
				} else {
					node.IsActive = false
				}
			}
		} else {
			// Node is completed with no next node to execute
			log.WithFields(log.Fields{
				"Id":  target.ID,
				"Ref": target.Ref,
			}).Info("Instance is finished")
			p.Active = false
		}
	} else {
		// Node is not completed
		log.WithFields(log.Fields{
			"Id":  target.ID,
			"Ref": target.Ref,
		}).Info("Wait for completion")
		return []*NodeInstance{target}
	}
	return make([]*NodeInstance, 0)
}

// NextNode find next node from edge
func (p *ProcessInstance) NextNode(ID string) (bool, *NodeInstance) {
	for _, edge := range p.Edges {
		if edge.Source == ID {
			found, target := p.FindNodeByID(edge.Target)
			return found && true, target
		}
	}
	return false, nil
}

// FindNodeByID find node by id
func (p *ProcessInstance) FindNodeByID(ID string) (bool, *NodeInstance) {
	for _, node := range p.Nodes {
		if node.ID == ID {
			return true, &node
		}
	}
	return false, nil
}

// CompleteByID complete node
func (p *ProcessInstance) CompleteByID(ID string, value bool) {
	for index, node := range p.Nodes {
		if node.ID == ID {
			p.Nodes[index].IsCompleted = value
			return
		}
	}
}
