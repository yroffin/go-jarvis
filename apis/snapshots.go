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
package apis

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	"github.com/yroffin/go-boot-sqllite/core/business"
	"github.com/yroffin/go-boot-sqllite/core/manager"
	"github.com/yroffin/go-boot-sqllite/core/models"
	app_models "github.com/yroffin/go-jarvis/models"
)

// Snapshot internal members
type Snapshot struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/snapshots"`
	// Swagger with injection mecanism
	Swagger core_apis.ISwaggerService `@autowired:"swagger"`
	// SwaggerService with injection mecanism
	Manager manager.IManager `@autowired:"manager"`
	// SqlCrudBusiness with injection mecanism
	SQLCrudBusiness business.ICrudBusiness `@autowired:"sql-crud-business"`
	// GraphBusiness with injection mecanism
	GraphBusiness business.ILinkBusiness `@autowired:"graph-crud-business"`
}

// ISnapshot implements IBean
type ISnapshot interface {
	core_apis.IAPI
}

// New constructor
func (p *Snapshot) New() ISnapshot {
	bean := Snapshot{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// SetSwagger inject notification
func (p *Snapshot) SetSwagger(value interface{}) {
	if assertion, ok := value.(core_apis.ISwaggerService); ok {
		p.Swagger = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetManager inject notification
func (p *Snapshot) SetManager(value interface{}) {
	if assertion, ok := value.(manager.IManager); ok {
		p.Manager = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetSQLCrudBusiness inject CrudBusiness
func (p *Snapshot) SetSQLCrudBusiness(value interface{}) {
	if assertion, ok := value.(business.ICrudBusiness); ok {
		p.SQLCrudBusiness = assertion
	} else {
		log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
	}
}

// SetGraphBusiness inject CrudBusiness
func (p *Snapshot) SetGraphBusiness(value interface{}) {
	if assertion, ok := value.(business.ILinkBusiness); ok {
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
	From int `json:"__from"`
	// To
	To int `json:"__to"`
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
			for idBean, entityBean := range entityBeanValue.(map[string]interface{}) {
				bean := p.Manager.GetBean(entityBeanType)
				if bean == nil {
					log.Println("Bean:", entityBeanType, "Not handled")
				} else {
					data, _ := json.MarshalIndent(entityBean, "", "\t")
					entity, _ := bean.(core_apis.CrudHandler).HandlerPost(string(data))
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
				var href = SnapshotHref{}
				json.Unmarshal(data, &href)
				if OldIDToNewID[strconv.Itoa(href.From)] != "" && OldIDToNewID[strconv.Itoa(href.To)] != "" {
					log.Println("Link:", NewIDToBean[OldIDToNewID[strconv.Itoa(href.From)]], "=[", linkedBeanType, "]=>", NewIDToBean[OldIDToNewID[strconv.Itoa(href.To)]])
					builds = append(builds, SnapshotHrefEntityBuild{
						Source: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[strconv.Itoa(href.From)]],
							ID:   OldIDToNewID[strconv.Itoa(href.From)],
						},
						Target: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[strconv.Itoa(href.To)]],
							ID:   OldIDToNewID[strconv.Itoa(href.To)],
						},
						IsError: false,
						Link:    linkedBeanType,
					})
				} else {
					if OldIDToNewID[strconv.Itoa(href.From)] == "" {
						hrefInErrors = append(hrefInErrors, href)
					}
					if OldIDToNewID[strconv.Itoa(href.To)] == "" {
						hrefInErrors = append(hrefInErrors, href)
					}
					builds = append(builds, SnapshotHrefEntityBuild{
						Source: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[strconv.Itoa(href.From)]],
							ID:   OldIDToNewID[strconv.Itoa(href.From)],
						},
						Target: SnapshotHrefEntity{
							Type: NewIDToBean[OldIDToNewID[strconv.Itoa(href.To)]],
							ID:   OldIDToNewID[strconv.Itoa(href.To)],
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
