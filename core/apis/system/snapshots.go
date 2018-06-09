// Package apis for common apis
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
package system

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/apis"
)

func init() {
	winter.Helper.Register("SnapshotBean", (&Snapshot{}).New())
}

// Snapshot internal members
type Snapshot struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/snapshots"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// SqlCrudBusiness with injection mecanism
	SQLCrudBusiness engine.ICrudBusiness `@autowired:"sql-crud-business"`
	// GraphBusiness with injection mecanism
	GraphBusiness engine.ILinkBusiness `@autowired:"graph-crud-business"`
}

// ISnapshot implements IBean
type ISnapshot interface {
	engine.IAPI
}

// New constructor
func (p *Snapshot) New() ISnapshot {
	bean := Snapshot{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Snapshot) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&SnapshotBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&SnapshotBeans{}).New()
	}
	p.HandlerTasks = func(name string, body string) (interface{}, int, error) {
		if name == "graph" {
			// task
			return p.GraphAll(body)
		}
		return "", -1, nil
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		if name == "restore" {
			// task
			return p.Restore(id, body)
		}
		if name == "download" {
			// task
			return p.Download(id, body)
		}
		return "", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Snapshot) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Snapshot) Validate(name string) error {
	return nil
}

// SnapshotHref simple Snapshot model
type SnapshotHref struct {
	// From
	From interface{} `json:"__from"`
	// To
	To interface{} `json:"__to"`
	// From
	FromStr string
	// To
	ToStr string
	// Href
	Href string `json:"href"`
	// Order
	Order string `json:"order"`
}

// SnapshotHrefEntity simple Snapshot model
type SnapshotHrefEntity struct {
	// Type
	Type string `json:"type"`
	// Id
	ID string `json:"id"`
}

// SnapshotHrefEntityBuild result
type SnapshotHrefEntityBuild struct {
	// Source
	Source SnapshotHrefEntity `json:"source"`
	// Target
	Target SnapshotHrefEntity `json:"Target"`
	// IsError
	IsError bool `json:"error"`
	// Link
	Link string `json:"link"`
	// Attr
	Attr map[string]string `json:"attributes"`
}

// Restore this Snapshot
func (p *Snapshot) Restore(id string, body string) (interface{}, int, error) {
	// Clear
	p.SQLCrudBusiness.Clear([]string{"SnapshotBean"})
	p.GraphBusiness.Clear()

	// Indexes
	var OldIDToNewID = make(map[string]string)
	var NewIDToOldID = make(map[string]string)
	var NewIDToBean = make(map[string]string)
	builds := make([]SnapshotHrefEntityBuild, 0)

	// retrieve command and serialize it
	model, _ := p.GetByID(id)
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Model")
	data := map[string]interface{}{}
	parsed := json.Unmarshal([]byte(model.(ISnapshotBean).GetJSON().(string)), &data)
	if parsed != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": parsed,
		}).Error("Model")
		return nil, -1, parsed
	}
	// iterate for entities
	for entityBeanType, entityBeanValue := range data {
		switch entityBeanType {
		case "HREF":
			break
		case "HREF_IF":
			break
		case "HREF_THEN":
			break
		default:
			log.WithFields(log.Fields{
				"bean": entityBeanType,
			}).Info("Restore")
			for idBean, entityBean := range entityBeanValue.(map[string]interface{}) {
				bean := winter.Helper.GetBean(entityBeanType)
				if bean == nil {
					log.WithFields(log.Fields{
						"bean": entityBeanType,
					}).Warn("Not handled")
				} else {
					data, _ := json.MarshalIndent(entityBean, "", "\t")
					entity, _ := bean.(engine.CrudHandler).HandlerPost(string(data))
					log.WithFields(log.Fields{
						"bean": entityBeanType,
						"with": entity.(models.IPersistent).GetID(),
					}).Info("Link")
					OldIDToNewID[idBean] = entity.(models.IPersistent).GetID()
					NewIDToOldID[entity.(models.IPersistent).GetID()] = idBean
					NewIDToBean[entity.(models.IPersistent).GetID()] = entityBeanType
				}
			}
			break
		}
	}
	hrefInErrors := make([]SnapshotHref, 0)
	// iterate for links
	for linkedBeanType, linkedBeanValue := range data {
		switch linkedBeanType {
		case "HREF", "HREF_IF", "HREF_THEN", "HREF_ELSE":
			for _, entityBean := range linkedBeanValue.(map[string]interface{}) {
				data, _ := json.MarshalIndent(entityBean, "", "\t")
				// Common href data
				var href = SnapshotHref{}
				json.Unmarshal(data, &href)
				assertFrom, okFrom := href.From.(string)
				if okFrom {
					href.FromStr = assertFrom
				} else {
					href.FromStr = strconv.Itoa(int(href.From.(float64)))
				}
				assertTo, okTo := href.To.(string)
				if okTo {
					href.ToStr = assertTo
				} else {
					href.ToStr = strconv.Itoa(int(href.To.(float64)))
				}
				// Attributes
				attr := make(map[string]string)
				for field, value := range entityBean.(map[string]interface{}) {
					assert, ok := value.(string)
					if ok {
						attr[field] = assert
					}
				}
				if OldIDToNewID[href.FromStr] != "" && OldIDToNewID[href.ToStr] != "" {
					log.WithFields(log.Fields{
						"from": NewIDToBean[OldIDToNewID[href.FromStr]],
						"with": linkedBeanType,
						"to":   NewIDToBean[OldIDToNewID[href.ToStr]],
					}).Info("Link")
					builds = append(builds, SnapshotHrefEntityBuild{
						Source: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[href.FromStr]],
							ID:   OldIDToNewID[href.FromStr],
						},
						Target: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[href.ToStr]],
							ID:   OldIDToNewID[href.ToStr],
						},
						IsError: false,
						Link:    linkedBeanType,
						Attr:    attr,
					})
				} else {
					if OldIDToNewID[href.FromStr] == "" {
						hrefInErrors = append(hrefInErrors, href)
					}
					if OldIDToNewID[href.ToStr] == "" {
						hrefInErrors = append(hrefInErrors, href)
					}
					builds = append(builds, SnapshotHrefEntityBuild{
						Source: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[href.FromStr]],
							ID:   OldIDToNewID[href.FromStr],
						},
						Target: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[href.ToStr]],
							ID:   OldIDToNewID[href.ToStr],
						},
						IsError: true,
					})
				}
			}
			break
		default:
			break
		}
	}
	// Create edge
	for _, edge := range builds {
		if !edge.IsError {
			toCreate := models.EdgeBean{
				Link:     edge.Link,
				Source:   edge.Source.Type,
				SourceID: edge.Source.ID,
				Target:   edge.Target.Type,
				TargetID: edge.Target.ID,
				Extended: make(map[string]interface{}),
			}
			for field, value := range edge.Attr {
				toCreate.Extended[field] = value
			}
			p.GraphBusiness.CreateLink(&toCreate)
		}
	}
	// HREF in errors
	for _, h := range hrefInErrors {
		log.WithFields(log.Fields{
			"from": h,
		}).Warn("No conversion")
	}
	return builds, len(builds), nil
}

// Download this Snapshot
func (p *Snapshot) Download(id string, body string) (interface{}, int, error) {
	output := make(map[string]interface{})

	out, _ := p.GraphBusiness.Export()
	for k, v := range out {
		log.WithFields(log.Fields{
			"link": k,
		}).Info("Link")
		index := make(map[string]interface{})
		for _, obj := range v {
			index[obj["id"].(string)] = obj
		}
		output[k] = index
	}

	beans := []string{"ProcessBean", "CommandBean", "DeviceBean", "ScriptPluginBean", "TriggerBean", "ViewBean", "ConnectorBean", "CronBean", "ConfigBean", "PropertyBean"}
	for _, b := range beans {
		log.WithFields(log.Fields{
			"name": b,
		}).Info("Bean")
		bean := winter.Helper.GetBean(b).(engine.IAPI)
		all, _ := bean.GetAll()
		index := make(map[string]interface{})
		for _, obj := range all {
			index[obj.GetID()] = obj
		}
		output[b] = index
	}

	return output, -1, nil
}

// Graph API
func (p *Snapshot) GraphAll(body string) (interface{}, int, error) {
	var options = []byte(`
		{
			"configure": {
				"enabled": true
			},
			"edges": {
				"smooth": {
					"type": "discrete",
					"forceDirection": "none",
					"roundness": 1
				}
			},
			"interaction": {
				"hover": true
			},
			"physics": {
				"forceAtlas2Based": {
					"gravitationalConstant": -26,
					"centralGravity": 0.005,
					"springLength": 230,
					"springConstant": 0.18
				},
				"solver": "forceAtlas2Based",
				"minVelocity": 0.75
			},
			"groups": {
				"DeviceBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf013",
						"size": 32,
						"color": "#57169a"
					}
				},
				"ScriptPluginBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf1e6",
						"size": 32,
						"color": "#57169a"
					}
				},
				"ViewBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf0e8",
						"size": 32,
						"color": "#57169a"
					}
				},
				"TriggerBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf0e7",
						"size": 32,
						"color": "#57169a"
					}
				},
				"CommandBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf02d",
						"size": 32,
						"color": "#57169a"
					}
				},
				"CronBean": {
					"shape": "icon",
					"icon": {
						"face": "FontAwesome",
						"code": "\uf133",
						"size": 32,
						"color": "#57169a"
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
	uniq := make(map[string]string)

	// Retrieve all nodes
	all, _ := p.GraphBusiness.All()
	for _, quad := range all {
		_, found := uniq[quad.SubjectID()]
		if !found {
			node := apis.Node{
				ID:    quad.SubjectID(),
				Label: quad.Subject(),
				Group: quad.Subject(),
				Title: quad.SubjectID(),
			}
			graph.Nodes = append(graph.Nodes, node)
			uniq[quad.SubjectID()] = quad.SubjectID()
		}
	}

	// Complete with targetID and sourceID
	for _, quad := range all {
		var data = make(map[string]interface{})
		json.Unmarshal([]byte(quad.Label()), &data)

		var targetID = data["targetId"].(string)
		_, foundTarget := uniq[targetID]
		if !foundTarget {
			node := apis.Node{
				ID:    targetID,
				Label: data["target"].(string),
				Group: data["target"].(string),
				Title: targetID,
			}
			graph.Nodes = append(graph.Nodes, node)
			uniq[targetID] = targetID
		}

		var sourceID = data["sourceId"].(string)
		_, foundSource := uniq[sourceID]
		if !foundSource {
			node := apis.Node{
				ID:    sourceID,
				Label: data["source"].(string),
				Group: data["source"].(string),
				Title: sourceID,
			}
			graph.Nodes = append(graph.Nodes, node)
			uniq[sourceID] = sourceID
		}
	}

	// Retrieve all edges
	for _, quad := range all {
		var data = map[string]interface{}{}
		json.Unmarshal([]byte(quad.Label()), &data)
		edge := apis.Edge{
			From:   quad.SubjectID(),
			To:     quad.ObjectID(),
			Label:  quad.Predicate(),
			Data:   quad.Label(),
			Title:  data["id"].(string),
			Smooth: true,
		}
		edge.Arrows = map[string]interface{}{}
		edge.Arrows["to"] = true
		graph.Edges = append(graph.Edges, edge)
	}
	return graph, -1, nil
}
