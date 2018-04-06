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
	"flag"
	"log"
	"os"
	"strings"

	"github.com/magiconair/properties"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("property-service", (&PropertyService{}).New())
}

// PropertyService internal members
type PropertyService struct {
	// members
	*winter.Service
	// Props
	props *properties.Properties
}

// IPropertyService implements IBean
type IPropertyService interface {
	winter.IBean
	Get(key string, def string) string
}

// New constructor
func (p *PropertyService) New() IPropertyService {
	bean := PropertyService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PropertyService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PropertyService) PostConstruct(name string) error {
	// init from a file
	p.props = properties.MustLoadFile("${HOME}/config.properties", properties.UTF8)
	// or from flags
	p.props.MustFlag(flag.CommandLine)
	flag.VisitAll(func(f *flag.Flag) {
		if strings.HasPrefix(f.Name, "D") {
			rune := []rune(f.Name)
			p.props.SetValue(string(rune[1:len(f.Name)]), f.Value.String())
		}
	})
	return nil
}

// Validate this SERVICE
func (p *PropertyService) Validate(name string) error {
	return nil
}

// Get value
func (p *PropertyService) Get(key string, def string) string {
	log.Println("Properties:", key, "Default:", def)
	value := p.props.GetString(key, def)
	if len(value) == 0 {
		fromEnv := os.Getenv(key)
		if len(fromEnv) == 0 {
			return def
		}
	}
	return value
}
