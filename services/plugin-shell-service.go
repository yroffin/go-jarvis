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
package services

import (
	"io/ioutil"
	"log"
	"os/exec"
	"reflect"

	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_apis "github.com/yroffin/go-boot-sqllite/core/services"
	app_models "github.com/yroffin/go-jarvis/models"
)

// PluginShellService internal members
type PluginShellService struct {
	// members
	*core_apis.SERVICE
	// SetPropertyService with injection mecanism
	SetPropertyService func(interface{}) `bean:"property-service"`
	PropertyService    *PropertyService
}

// IPluginShellService implements IBean
type IPluginShellService interface {
	core_bean.IBean
}

// Init this SERVICE
func (p *PluginShellService) Init() error {
	// inject store
	p.SetPropertyService = func(value interface{}) {
		if assertion, ok := value.(*PropertyService); ok {
			p.PropertyService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
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
func (p *PluginShellService) Call(body string) (app_models.ValueBean, error) {
	cmd := exec.Command("sh", "-c", body)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	capture := app_models.ValueBean{}

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