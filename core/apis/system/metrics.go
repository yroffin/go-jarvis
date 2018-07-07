// Package system for metrics
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
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services/mqtt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	winter.Helper.Register("MetrictsBean", (&Metrics{}).New())
}

// Metrics internal members
type Metrics struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// Api
	ProHandler interface{} `@handler:"Handler" path:"/api/metrics" method:"GET" mime-type:"/application/json"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// Internal
	counters map[string]prometheus.Counter
	gauges   map[string]prometheus.Gauge
}

// IMetrics implements IBean
type IMetrics interface {
	engine.IAPI
	AddCounter(name string, help string)
	SetCounter(name string, value float64)
	IncCounter(name string)
	AddGauge(name string, help string)
	SetGauge(name string, value float64)
	IncGauge(name string)
}

// New constructor
func (p *Metrics) New() IMetrics {
	bean := Metrics{API: &engine.API{Bean: &winter.Bean{}}}
	bean.counters = map[string]prometheus.Counter{}
	bean.gauges = map[string]prometheus.Gauge{}
	return &bean
}

// Init this API
func (p *Metrics) Init() error {
	return p.API.Init()
}

// PostConstruct this API
func (p *Metrics) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)

	// Metrics
	p.AddCounter("cpu_user", "Current user cpu time.")
	p.AddCounter("cpu_idle", "Current idle cpu time.")
	p.AddCounter("cpu_system", "Current system cpu time.")

	// Start metrics
	go p.SystemMetrics()
	return nil
}

// Validate this API
func (p *Metrics) Validate(name string) error {
	return nil
}

// SystemMetrics API
func (p *Metrics) SystemMetrics() func() error {
	for {
		m, _ := mem.VirtualMemory()
		c, _ := cpu.Times(false)
		// Set counter
		p.SetCounter("cpu_idle", c[0].Idle)
		p.SetCounter("cpu_user", c[0].User)
		p.SetCounter("cpu_system", c[0].System)
		// Notify system ready
		p.MqttService.PublishMostOne("/system/metrics/memory", models.ToJSON(m))
		p.MqttService.PublishMostOne("/system/metrics/cpu", models.ToJSON(c))
		time.Sleep(5 * time.Second)
	}
}

// Handler handler
func (p *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

// AddCounter add prometheus counter
func (p *Metrics) AddCounter(name string, help string) {
	_, exist := p.counters[name]
	if !exist {
		p.counters[name] = prometheus.NewCounter(prometheus.CounterOpts{
			Name: name,
			Help: help,
		})
		err := prometheus.Register(p.counters[name])
		if err == nil {
			return
		}
		log.WithFields(log.Fields{
			"Error": err,
		}).Warn("Prometheus register counter")
		return
	}
	log.WithFields(log.Fields{
		"Name": name,
	}).Warn("Prometheus register counter already done")
}

// SetCounter set prometheus counter
func (p *Metrics) SetCounter(name string, value float64) {
	found, exist := p.counters[name]
	if exist {
		found.Set(value)
		p.MqttService.PublishMostOne("/system/metrics/counters/"+name, models.ToJSON(value))
	}
}

// IncCounter inc prometheus counter
func (p *Metrics) IncCounter(name string) {
	found, exist := p.counters[name]
	if exist {
		found.Inc()
		p.MqttService.PublishMostOne("/system/metrics/counters/"+name, models.ToJSON(-1))
	}
}

// AddGauge add prometheus gauge
func (p *Metrics) AddGauge(name string, help string) {
	_, exist := p.counters[name]
	if !exist {
		p.gauges[name] = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		})
		err := prometheus.Register(p.gauges[name])
		if err == nil {
			return
		}
		log.WithFields(log.Fields{
			"Name":  name,
			"Error": err,
		}).Warn("Prometheus register gauge")
		return
	}
	log.WithFields(log.Fields{
		"Name": name,
	}).Warn("Prometheus register gauge already done")
}

// SetGauge set prometheus gauge
func (p *Metrics) SetGauge(name string, value float64) {
	found, exist := p.gauges[name]
	if exist {
		found.Set(value)
		p.MqttService.PublishMostOne("/system/metrics/gauges/"+name, models.ToJSON(value))
	}
}

// IncGauge inc prometheus gauge
func (p *Metrics) IncGauge(name string) {
	found, exist := p.gauges[name]
	if exist {
		found.Inc()
		p.MqttService.PublishMostOne("/system/metrics/gauges/"+name, models.ToJSON(-1))
	}
}
