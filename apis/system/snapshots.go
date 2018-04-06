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
	"log"
	"reflect"
	"strconv"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	app_models "github.com/yroffin/go-jarvis/models"
)

func init() {
	engine.Winter.Register("SnapshotBean", (&Snapshot{}).New())
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
	// SwaggerService with injection mecanism
	Manager engine.IManager `@autowired:"manager"`
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
	bean := Snapshot{API: &engine.API{Bean: &engine.Bean{}}}
	return &bean
}

// SetSwagger inject notification
func (p *Snapshot) SetSwagger(value interface{}) {
	if assertion, ok := value.(engine.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetManager inject notification
func (p *Snapshot) SetManager(value interface{}) {
	if assertion, ok := value.(engine.IManager); ok {
		p.Manager = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetSQLCrudBusiness inject CrudBusiness
func (p *Snapshot) SetSQLCrudBusiness(value interface{}) {
	if assertion, ok := value.(engine.ICrudBusiness); ok {
		p.SQLCrudBusiness = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetGraphBusiness inject CrudBusiness
func (p *Snapshot) SetGraphBusiness(value interface{}) {
	if assertion, ok := value.(engine.ILinkBusiness); ok {
		p.GraphBusiness = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// Init this API
func (p *Snapshot) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&app_models.SnapshotBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&app_models.SnapshotBeans{}).New()
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
	model := (&app_models.SnapshotBean{}).New()
	p.GetByID(id, model)
	// iterate for entities
	for entityBeanType, entityBeanValue := range model.GetJSON().(map[string]interface{}) {
		switch entityBeanType {
		case "HREF":
			break
		case "HREF_IF":
			break
		case "HREF_THEN":
			break
		default:
			log.Println("Bean:", entityBeanType)
			for idBean, entityBean := range entityBeanValue.(map[string]interface{}) {
				bean := p.Manager.GetBean(entityBeanType)
				if bean == nil {
					log.Println("Bean:", entityBeanType, "Not handled")
				} else {
					data, _ := json.MarshalIndent(entityBean, "", "\t")
					entity, _ := bean.(engine.CrudHandler).HandlerPost(string(data))
					log.Println("Type:", entityBeanType, "With:", entity.(models.IPersistent).GetID())
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
	for linkedBeanType, linkedBeanValue := range model.GetJSON().(map[string]interface{}) {
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
				assertTo, okTo := href.From.(string)
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
					log.Println("Link:", NewIDToBean[OldIDToNewID[href.FromStr]], "=[", linkedBeanType, "]=>", NewIDToBean[OldIDToNewID[href.ToStr]])
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
		log.Println("Warn:", h, "conversion is null")
	}
	return builds, len(builds), nil
}

// Download this Snapshot
func (p *Snapshot) Download(id string, body string) (interface{}, int, error) {
	output := make(map[string]interface{})

	out, _ := p.GraphBusiness.Export()
	for k, v := range out {
		log.Println("Link:", k)
		index := make(map[string]interface{})
		for _, obj := range v {
			index[obj["id"].(string)] = obj
		}
		output[k] = index
	}

	beans := []string{"ProcessBean", "CommandBean", "DeviceBean", "ScriptPluginBean", "TriggerBean", "ViewBean", "ConnectorBean", "CronBean", "ConfigBean", "PropertyBean"}
	for _, b := range beans {
		log.Println("Bean:", b)
		bean := p.Manager.GetBean(b).(engine.IAPI)
		all, _ := bean.GetAll()
		index := make(map[string]interface{})
		for _, obj := range all {
			index[obj.GetID()] = obj
		}
		output[b] = index
	}

	return output, -1, nil
}
