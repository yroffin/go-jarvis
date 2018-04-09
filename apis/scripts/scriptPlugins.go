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
	"log"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/apis/commands"
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
	Execute(id string, parameters map[string]interface{}) (interface{}, int, error)
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
			return p.Execute(id, parameters)
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

// Execute the script
func (p *ScriptPlugin) Execute(id string, parameters map[string]interface{}) (interface{}, int, error) {
	// Retrieve all script
	links, _ := p.GetAllLinks(id, p.LinkCommand)
	for _, command := range links {
		log.Println("COMMAND - INPUT", command.GetID(), parameters)
		result, count, _ := p.LinkCommand.Execute(command.GetID(), parameters)
		log.Println("COMMAND - OUTPUT", command.GetID(), result, count)
	}
	return links, -1, nil
}

// Graph type
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// Node type
type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Edge type
type Edge struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Label string `json:"label"`
}

// Graph render as graph
func (p *ScriptPlugin) Graph(id string, parameters map[string]interface{}) (interface{}, int, error) {
	graph := Graph{
		Nodes: make([]Node, 0),
		Edges: make([]Edge, 0),
	}
	// Retrieve all script
	instance, _ := p.GetByID(id)
	graph.Nodes = append(graph.Nodes, Node{ID: instance.GetID(), Label: instance.(IScriptPluginBean).GetName()})
	links, _ := p.GetAllLinks(id, p.LinkCommand)
	for _, command := range links {
		graph.Nodes = append(graph.Nodes, Node{ID: command.GetID(), Label: command.(commands.ICommandBean).GetName()})
		graph.Edges = append(graph.Edges, Edge{From: instance.GetID(), To: command.GetID(), Label: "link"})
	}
	return graph, -1, nil
}
