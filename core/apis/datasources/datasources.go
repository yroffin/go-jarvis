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
package datasources

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/helpers"
	"github.com/yroffin/go-jarvis/core/services"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
)

func init() {
	winter.Helper.Register("DataSourceBean", (&DataSource{}).New())
}

// DataSource internal members
type DataSource struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// mounts
	Crud interface{} `@crud:"/api/datasources"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// SetPropertyService with injection mecanism
	PropertyService services.IPropertyService `@autowired:"property-service"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
	// Prometheus
	prometheus *helpers.HTTPClient
}

// IDataSource implements IBean
type IDataSource interface {
	engine.IAPI
}

// New constructor
func (p *DataSource) New() IDataSource {
	bean := DataSource{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *DataSource) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&DataSourceBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&DataSourceBeans{}).New()
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *DataSource) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	// Init the client
	p.prometheus = &helpers.HTTPClient{URL: p.PropertyService.Get("jarvis.prometheus.url", "http://192.168.1.26:9090")}
	return nil
}

// Validate this API
func (p *DataSource) Validate(name string) error {
	// Start metrics
	go p.Discover()

	return nil
}

// Discover datasources
func (p *DataSource) Discover() func() error {
	for {
		headers := make(map[string]string)
		params := make(map[string]string)
		params["query"] = "up"
		result, err := p.prometheus.GET(p.PropertyService.Get("jarvis.prometheus.api.query", "/api/v1/query"), headers, params)

		if err == nil {
			// Unmarshall result
			var data = result["data"].(map[string]interface{})
			if data["resultType"].(string) == "vector" {
				var results = data["result"].([]interface{})
				for _, result := range results {
					var metric = result.(map[string]interface{})["metric"].(map[string]interface{})
					var ref = metric["job"].(string) + "@" + metric["instance"].(string)
					var all, _ = p.GetAll()
					var found = false
					for _, datasource := range all {
						log.WithFields(log.Fields{
							"metric":     metric,
							"datasource": datasource,
						}).Warn("Results")
						if datasource.(*DataSourceBean).ExternalRef == ref {
							found = true
						}
					}
					if !found {
						toCreate := DataSourceBean{
							Name:        metric["job"].(string),
							Icon:        "table",
							ExternalRef: ref,
						}
						p.SQLCrudBusiness.Create(&toCreate)
					}
				}
				// Notify system ready
				p.MqttService.PublishMostOne("/system/datasources/discover", models.ToJSON(result))
			}
		} else {
			log.WithFields(log.Fields{
				"url":   p.prometheus.URL,
				"error": err,
			}).Warn("Prometheus")
		}
		time.Sleep(10 * time.Second)
	}
}
