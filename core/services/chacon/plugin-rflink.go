// Package chacon for common services
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
package chacon

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
	serial "go.bug.st/serial.v1"
)

func init() {
	winter.Helper.Register("plugin-rflink-service", (&PluginRFLinkService{}).New())
}

// PluginRFLinkService internal members
type PluginRFLinkService struct {
	// members
	*winter.Service
	// PropertyService with injection mecanism
	PropertyService services.IPropertyService `@autowired:"property-service"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
	// Port
	handle serial.Port
	// Chan
	channel chan string
}

// IPluginRFLinkService implements IBean
type IPluginRFLinkService interface {
	// Extend bean
	winter.IService
	// Local method
	Chacon(channel string, command string, order string) (models.IValueBean, error)
}

// New constructor
func (p *PluginRFLinkService) New() IPluginRFLinkService {
	bean := PluginRFLinkService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginRFLinkService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PluginRFLinkService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginRFLinkService) Validate(name string) error {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Rflink - get port list")
	}
	if len(ports) == 0 {
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Rflink - no port found")
		}
	}
	log.WithFields(log.Fields{
		"ports": ports,
	}).Info("Rflink - get port list")
	errOpen := p.Open()
	if errOpen == nil {
		// go p.Read()
		// go p.Write()
		return nil
	}
	// Notify system ready
	p.MqttService.PublishMostOne("/system/rflink", "ready")
	return err
}

// Open init
func (p *PluginRFLinkService) Open() error {
	comport := p.PropertyService.Get("jarvis.rflink.comport", "/dev/ttyS0")
	bitRate, _ := strconv.Atoi(p.PropertyService.Get("jarvis.rflink.baud", "57600"))
	log.WithFields(log.Fields{
		"key":  comport,
		"baud": bitRate,
	}).Info("Rflink - comport open")
	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: bitRate,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(comport, mode)
	if err != nil {
		log.WithFields(log.Fields{
			"options": mode,
			"err":     err,
		}).Error("Rflink - comport open")
	}

	p.handle = port

	// Create channel
	p.channel = make(chan string)
	return nil
}

// Read serial port
func (p *PluginRFLinkService) Read() error {
	for {
		buf := make([]byte, 128)
		n, err := p.handle.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		p.channel <- string(buf[:n])
		log.WithFields(log.Fields{
			"value": string(buf[:n]),
		}).Info("Rflink - com read")
	}
}

// Read serial port
func (p *PluginRFLinkService) Write() error {
	for {
		value := <-p.channel
		log.WithFields(log.Fields{
			"value": value,
		}).Info("Rflink - com write")
	}
}

// Chacon execution
func (p *PluginRFLinkService) Chacon(channel string, command string, order string) (models.IValueBean, error) {
	result := (&models.ValueBean{}).New()
	result.SetString("Channel", channel)
	result.SetString("Command", command)
	result.SetString("Order", order)
	return result, nil
}
