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
package scripts

import (
	"encoding/json"
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/apis"
	"github.com/yroffin/go-jarvis/core/apis/commands"
)

func init() {
	winter.Helper.Register("ScriptPluginBean", (&ScriptPlugin{}).New())
}

// ScriptPlugin internal members
type ScriptPlugin struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/plugins/scripts"`
	// LinkCommand with injection mecanism
	LinkCommand commands.ICommand `@autowired:"CommandBean" @link:"/api/plugins/scripts" @href:"commands"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// IScriptPlugin implements IBean
type IScriptPlugin interface {
	engine.IAPI
	RenderOrExecute(id string, parameters map[string]interface{}, execute bool) (interface{}, int, error)
}

// New constructor
func (p *ScriptPlugin) New() IScriptPlugin {
	bean := ScriptPlugin{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *ScriptPlugin) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&ScriptPluginBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&ScriptPluginBeans{}).New()
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		var parameters = make(map[string]interface{})
		json.Unmarshal([]byte(body), &parameters)
		if name == "execute" {
			// Execute the script
			return p.RenderOrExecute(id, parameters, true)
		}
		if name == "render" {
			// Execute the script
			return p.RenderOrExecute(id, parameters, false)
		}
		if name == "graph" {
			// Execute the script
			return p.Graph(id, parameters)
		}
		return "{}", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *ScriptPlugin) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *ScriptPlugin) Validate(name string) error {
	return nil
}

// RenderOrExecute the script
func (p *ScriptPlugin) RenderOrExecute(id string, parameters map[string]interface{}, execute bool) (interface{}, int, error) {
	// Store outputs
	outputs := make(map[string]interface{})
	// Retrieve all script
	links, _ := p.GetAllLinks(id, p.LinkCommand)
	// Sort
	sort.Slice(links, func(i, j int) bool {
		var left = links[i].(commands.ICommandBean).GetExtend()["order"]
		var right = links[j].(commands.ICommandBean).GetExtend()["order"]
		iLeft, _ := strconv.Atoi(left.(string))
		iRight, _ := strconv.Atoi(right.(string))
		return iLeft < iRight
	})
	log.WithFields(log.Fields{
		"script":     id,
		"parameters": parameters,
	}).Info("Render or execute script")
	// Iterate
	for _, command := range links {
		var t = command.(commands.ICommandBean).GetExtend()["type"]

		// ignore data in phase action
		if t == "data" && execute {
			log.WithFields(log.Fields{
				"type":       t,
				"script.id":  id,
				"command.id": command.GetID(),
			}).Warn("Cannot be executed, it's a data")
			continue
		}
		// ignore action in phase data
		if t == "action" && !execute {
			log.WithFields(log.Fields{
				"type":       t,
				"script.id":  id,
				"command.id": command.GetID(),
			}).Warn("Cannot be rendered, it's a action")
			continue
		}

		// Execute the command in all case to obtain data
		// to render ... or simply execute
		var name = command.(commands.ICommandBean).GetExtend()["name"]
		result, _, _ := p.LinkCommand.Execute(command.GetID(), parameters)

		// store result
		if name != nil {
			log.WithFields(log.Fields{
				"name":   name,
				"result": result,
			}).Info("Render or execute script - result")
			outputs[name.(string)] = result
		} else {
			log.WithFields(log.Fields{
				"name": nil,
			}).Warn("Store result cannot be done")
		}
	}
	log.WithFields(log.Fields{
		"outputs": models.ToJSON(outputs),
	}).Debug("Render or execute script - result")
	return outputs, len(outputs), nil
}

// Graph render as graph
func (p *ScriptPlugin) Graph(id string, parameters map[string]interface{}) (interface{}, int, error) {
	graph := apis.Graph{
		Nodes: make([]apis.Node, 0),
		Edges: make([]apis.Edge, 0),
	}
	// Retrieve all script
	instance, _ := p.GetByID(id)
	graph.Nodes = append(graph.Nodes, apis.Node{ID: instance.GetID(), Label: instance.(IScriptPluginBean).GetName()})
	links, _ := p.GetAllLinks(id, p.LinkCommand)
	for _, command := range links {
		graph.Nodes = append(graph.Nodes, apis.Node{ID: command.GetID(), Label: command.(commands.ICommandBean).GetName()})
		graph.Edges = append(graph.Edges, apis.Edge{From: instance.GetID(), To: command.GetID(), Label: "link"})
	}
	return graph, -1, nil
}
