// Package mqtt for common interfaces
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
package teleinfo

import (
	"bufio"
	"os"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/core/services"
	"github.com/yroffin/go-jarvis/core/services/mqtt"
)

func init() {
	winter.Helper.Register("teleinfo-service", (&service{}).New())
}

// On linux
// apply: stty 1200 cs7 evenp cstopb -igncr -inlcr -brkint -icrnl -opost -isig -icanon -iexten -F /dev/ttyUSB0
// to configure teleinfo tty
// Cf. https://hallard.me/gestion-de-la-teleinfo-avec-un-raspberry-pi-et-une-carte-arduipi/

// service internal members
type service struct {
	// members
	*winter.Service
	// SetPropertyService with injection mecanism
	PropertyService services.IPropertyService `@autowired:"property-service"`
	// IMqttService with injection mecanism
	MqttService mqtt.IMqttService `@autowired:"mqtt-service"`
	// internal
	stream  chan byte
	entries map[string]string
	mutex   *sync.Mutex
}

// IMqttService implements IBean
type IMqttService interface {
	winter.IService
}

// New constructor
func (p *service) New() IMqttService {
	bean := service{Service: &winter.Service{Bean: &winter.Bean{}}}
	bean.stream = make(chan byte, 1024)
	bean.entries = make(map[string]string)
	bean.mutex = &sync.Mutex{}
	return &bean
}

// Init this API
func (p *service) Init() error {
	return nil
}

// PostConstruct this API
func (p *service) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *service) Validate(name string) error {
	// Activate service only if opion is set
	if len(p.PropertyService.MustGet("jarvis.option.teleinfo.file")) > 0 {
		// start worker
		go p.handleReadFile(p.PropertyService.MustGet("jarvis.option.teleinfo.file"))
		go p.worker()

		// log information
		log.WithFields(log.Fields{
			"teleinfoFile": p.PropertyService.MustGet("jarvis.option.teleinfo.file"),
		}).Info("teleinfo")
	}
	return nil
}

func (p *service) handleReadFile(device string) error {

	s, err := os.OpenFile(device, syscall.O_RDONLY, 0666)

	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("teleinfo")
	}

	log.WithFields(log.Fields{
		"Device": device,
	}).Info("teleinfo")

	buffer := make([]byte, 4096)
	reader := bufio.NewReader(s)

	// Receive reply
	for {
		if len, err := reader.Read(buffer); err != nil {
			// sleep while no bytes
			// to avoid system flood read
			log.WithFields(log.Fields{
				"Error": err,
			}).Error("teleinfo")
			time.Sleep(1000 * time.Millisecond)
		} else {
			// dispatch io
			for i := 0; i < len; i++ {
				p.stream <- buffer[i]
			}
			time.Sleep(2000 * time.Millisecond)
		}
	}
}

// Trame trame EDF
type Trame struct {
	etiquette string // ETIQUETTE (4 à 8 caractères)
	data      string // DATA (1 à 12 caractères)
	checksum  string // CHECKSUM (caractère de contrôle : Somme de caractère)
	line      string // all bytes
	sum       int    // CHECKSUM calculé
}

/**
* phase1: LF (0x0A)
* phase2: ETIQUETTE (4 à 8 caractères)
* phase3: SP (0x20)
* phase4: DATA (1 à 12 caractères)
* phase5: SP (0x20)
* phase6: CHECKSUM (caractère de contrôle : Somme de caractère)
* CR (0x0D)
 */

// submit trame
func (p *service) submit(trame Trame) {
	p.mutex.Lock()
	p.entries[trame.etiquette] = trame.data
	// Notify system ready
	p.MqttService.PublishMostOne("/system/teleinfo", trame.data)
	p.mutex.Unlock()
}

func (p *service) handleTrame(trame string) {
	var espace int
	var send Trame
	for i := 0; i < len(trame); i++ {
		switch {
		case trame[i] == 0x20:
			espace++
			send.line += string([]byte{trame[i]})
			continue
		default:
			if espace == 0 {
				send.etiquette += string([]byte{trame[i]})
				send.line += string([]byte{trame[i]})
			}
			if espace == 1 {
				send.data += string([]byte{trame[i]})
				send.line += string([]byte{trame[i]})
			}
			if espace == 2 {
				send.checksum += string([]byte{trame[i]})
			}
			continue
		}
	}
	send.sum = 0
	for i := 0; i < len(send.line)-1; i++ {
		send.sum += int(send.line[i])
	}
	// submit new value
	// Cf. http://forum.arduino.cc/index.php?topic=300157.0 pour le checksum
	// send when sum is ok
	var iCheck = 0
	// sometimes checksum is null
	if len(send.checksum) > 0 {
		iCheck = int(send.checksum[0])
	} else {
		iCheck = 0
	}
	if ((send.sum & 0x3F) + 0x20) == iCheck {
		p.submit(send)
	} else {
		log.Error("teleinfo", log.Fields{
			"submit":                send.etiquette,
			"data":                  send.data,
			"checksum":              send.checksum,
			"checksum/int":          iCheck,
			"line":                  send.line,
			"checksum/computed":     string((send.sum & 0x3F) + 0x20),
			"checksum/computed/int": (send.sum & 0x3F) + 0x20,
			"trame":                 trame,
		})
	}
}

func (p *service) handleTrames(trame string) {
	var send string
	for i := 0; i < len(trame); i++ {
		switch {
		case trame[i] == 0x0A:
			send = ""
			continue
		case trame[i] == 0x0D:
			p.handleTrame(send)
			continue
		default:
			send += string([]byte{trame[i]})
		}
	}
}

// worker to consume file
func (p *service) worker() {
	var trame string
	var etx bool

	// wait for ETX 0x003
	for i := 0; etx == false; i++ {
		var x = <-p.stream
		if x == 0x03 {
			etx = true
		}
	}

	// daemon
	for {
		var x = <-p.stream
		switch {
		case x == 0x03:
			// wait for ETX 0x003
			p.handleTrames(trame)
			break
		case x == 0x02:
			// wait for STX 0x002
			trame = ""
			break
		case x != 0x02 && x != 0x03:
			// other
			trame += string([]byte{x})
			break
		}
	}
}

// GetEntries load entries
func (p *service) GetEntries(entries map[string]string) map[string]string {
	p.mutex.Lock()
	for key, value := range p.entries {
		entries[key] = value
	}
	p.mutex.Unlock()
	return entries
}
