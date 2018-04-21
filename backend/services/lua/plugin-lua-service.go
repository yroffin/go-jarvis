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
package lua

import (
	"net/http"
	"reflect"

	lua_http "github.com/cjoudrey/gluahttp"
	log "github.com/sirupsen/logrus"
	lua_mapper "github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"

	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	app_services "github.com/yroffin/go-jarvis/services"
)

func init() {
	winter.Helper.Register("plugin-lua-service", (&PluginLuaService{}).New())
}

// PluginLuaService internal members
type PluginLuaService struct {
	// members
	*winter.Service
	// SetPropertyService with injection mecanism
	PropertyService app_services.IPropertyService `@autowired:"property-service"`
}

// IPluginLuaService implements IBean
type IPluginLuaService interface {
	// Extend bean
	winter.IBean
	// Local method
	Call(body string, args map[string]interface{}) (models.IValueBean, error)
}

// New constructor
func (p *PluginLuaService) New() IPluginLuaService {
	bean := PluginLuaService{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this SERVICE
func (p *PluginLuaService) Init() error {
	return nil
}

// PostConstruct this SERVICE
func (p *PluginLuaService) PostConstruct(name string) error {
	return nil
}

// Validate this SERVICE
func (p *PluginLuaService) Validate(name string) error {
	return nil
}

// Call execution
func (p *PluginLuaService) Call(body string, args map[string]interface{}) (models.IValueBean, error) {
	l := lua.NewState()
	defer l.Close()

	l.PreloadModule("http", lua_http.NewHttpModule(&http.Client{}).Loader)
	l.PreloadModule("extends", Loader)

	// register functions to the table
	mod := l.NewTable()

	l.SetGlobal("input", mod)
	for k, v := range args {
		switch reflect.TypeOf(v).String() {
		case "string":
			l.SetField(mod, k, lua.LString(v.(string)))
			break
		case "float64":
			l.SetField(mod, k, lua.LNumber(v.(float64)))
			break
		default:
			log.WithFields(log.Fields{
				"type": reflect.TypeOf(v).String(),
			}).Error("Unable to map type")
			break
		}
	}

	result := (&models.ValueBean{}).New()

	if err := l.DoString(body); err != nil {
		result.Set("error", err)
		return result, nil
	}

	// Map lua table to result
	var res map[string]interface{}
	if err := lua_mapper.Map(l.Get(-1).(*lua.LTable), &res); err != nil {
		result.Set("error", err)
		return result, nil
	}

	// build result
	for k, v := range res {
		result.Set(k, v)
	}

	return result, nil
}
