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
	user   prometheus.Counter
	idle   prometheus.Counter
	system prometheus.Counter
}

// IMetrics implements IBean
type IMetrics interface {
	engine.IAPI
}

// New constructor
func (p *Metrics) New() IMetrics {
	bean := Metrics{API: &engine.API{Bean: &winter.Bean{}}}
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
	p.user = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cpu_user",
		Help: "Current user cpu time.",
	})
	p.idle = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cpu_idle",
		Help: "Current idle cpu time.",
	})
	p.system = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cpu_system",
		Help: "Current system cpu time.",
	})
	// Metrics have to be registered to be exposed
	prometheus.MustRegister(p.user)
	prometheus.MustRegister(p.idle)
	prometheus.MustRegister(p.system)

	return nil
}

// Validate this API
func (p *Metrics) Validate(name string) error {
	// Start metrics
	go p.SystemMetrics()

	return nil
}

// SystemMetrics API
func (p *Metrics) SystemMetrics() func() error {
	for {
		m, _ := mem.VirtualMemory()
		c, _ := cpu.Times(false)
		p.idle.Set(c[0].Idle)
		p.user.Set(c[0].User)
		p.system.Set(c[0].System)
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
