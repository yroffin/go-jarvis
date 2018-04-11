// Package services for common services
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
package shell

import (
	"io/ioutil"
	"os/exec"

	log "github.com/sirupsen/logrus"

	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	app_services "github.com/yroffin/go-jarvis/services"
)

func init() {
	winter.Helper.Register("plugin-shell-service", (&PluginShellService{}).New())
}

// PluginShellService internal members
type PluginShellService struct {
	// members
	*winter.Service
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginShellService implements IBean
type IPluginShellService interface {
	winter.IBean
	Call(body string) (core_models.IValueBean, error)
}

// New constructor
func (p *PluginShellService) New() IPluginShellService {
	bean := PluginShellService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginShellService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PluginShellService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginShellService) Validate(name string) error {
	return nil
}

func build(data []byte) []string {
	var lines = 0
	for i := 1; i < len(data); i++ {
		switch data[i] {
		case 10, 13:
			lines++
			break
		}
	}

	var accu = ""
	console := make([]string, lines)
	var index = 0
	var byt = 0
	for index < lines {
		switch data[byt] {
		case 13:
			byt++
			break
		case 10:
			console[index] = accu
			index++
			byt++
			accu = ""
			break
		default:
			accu += string(data[byt])
			byt++
		}
	}

	return console
}

// Call execution
func (p *PluginShellService) Call(body string) (core_models.IValueBean, error) {
	cmd := exec.Command("sh", "-c", body)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Stdout")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Stderr")
	}

	if err := cmd.Start(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Start")
	}

	capture := (&core_models.ValueBean{}).New()

	slurpOut, _ := ioutil.ReadAll(stdout)
	capture.Set("stdout", build(slurpOut))

	slurpErr, _ := ioutil.ReadAll(stderr)
	capture.Set("stderr", build(slurpErr))

	if err := cmd.Wait(); err != nil {
		capture.Set("exit", err.Error())
		capture.Set("code", -1)
		return capture, nil
	}

	capture.Set("exit", "No error")
	capture.Set("code", 0)
	return capture, nil
}
