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
package apis

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	app_models "github.com/yroffin/go-jarvis/models"
	app_chacon "github.com/yroffin/go-jarvis/services/chacon"
	app_lua "github.com/yroffin/go-jarvis/services/lua"
	app_shell "github.com/yroffin/go-jarvis/services/shell"
	app_slack "github.com/yroffin/go-jarvis/services/slack"
	app_zway "github.com/yroffin/go-jarvis/services/zway"
)

// Command internal members
type Command struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	crud string `path:"/api/commands"`
	// SlackService with injection mecanism
	SetSlackService func(interface{}) `bean:"slack-service"`
	SlackService    *app_slack.SlackService
	// ShellService with injection mecanism
	SetShellService func(interface{}) `bean:"shell-service"`
	ShellService    *app_shell.ShellService
	// LuaService with injection mecanism
	SetLuaService func(interface{}) `bean:"lua-service"`
	LuaService    *app_lua.LuaService
	// ChaconService with injection mecanism
	SetChaconService func(interface{}) `bean:"chacon-service"`
	ChaconService    *app_chacon.ChaconService
	// ZwayService with injection mecanism
	SetZwayService func(interface{}) `bean:"zway-service"`
	ZwayService    *app_zway.ZwayService
}

// ICommand implements IBean
type ICommand interface {
	core_bean.IBean
}

// Init this API
func (p *Command) Init() error {
	// inject store
	p.SetSlackService = func(value interface{}) {
		if assertion, ok := value.(*app_slack.SlackService); ok {
			p.SlackService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// inject store
	p.SetShellService = func(value interface{}) {
		if assertion, ok := value.(*app_shell.ShellService); ok {
			p.ShellService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// inject store
	p.SetLuaService = func(value interface{}) {
		if assertion, ok := value.(*app_lua.LuaService); ok {
			p.LuaService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// inject store
	p.SetChaconService = func(value interface{}) {
		if assertion, ok := value.(*app_chacon.ChaconService); ok {
			p.ChaconService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// inject store
	p.SetZwayService = func(value interface{}) {
		if assertion, ok := value.(*app_zway.ZwayService); ok {
			p.ZwayService = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// Crud
	p.HandlerGetAll = func() (string, error) {
		return p.GenericGetAll((&app_models.CommandBean{}).New(), (&app_models.CommandBeans{}).New())
	}
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GenericGetByID(id, (&app_models.CommandBean{}).New())
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.GenericPost(body, (&app_models.CommandBean{}).New())
	}
	p.HandlerTasks = func(name string, body string) (string, error) {
		return "", nil
	}
	p.HandlerTasksByID = func(id string, name string, body string) (string, error) {
		if name == "execute" {
			// task
			return p.Execute(id, body)
		}
		if name == "test" {
			// task
			res, err := p.Test(id, body)
			return strconv.FormatBool(res), err
		}
		return "", nil
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.GenericPutByID(id, body, (&app_models.CommandBean{}).New())
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.GenericDeleteByID(id, (&app_models.CommandBean{}).New())
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.GenericPatchByID(id, body, (&app_models.CommandBean{}).New())
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Command) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p)
	return nil
}

// Validate this API
func (p *Command) Validate(name string) error {
	return nil
}

// New constructor
func (p *Command) New() ICommand {
	bean := Command{API: &core_apis.API{Bean: &core_bean.Bean{}}}
	return &bean
}

// Execute this command
func (p *Command) decode(id string, body string) (string, app_models.ICommandBean, map[string]interface{}, error) {
	// retrieve command and serialize it
	model := (&app_models.CommandBean{}).New()
	p.GetByID(id, model)
	raw, _ := json.Marshal(&model)
	converted := make(map[string]interface{})
	json.Unmarshal(raw, &converted)
	// retrieve args and serialize it
	args := make(map[string]interface{})
	json.Unmarshal([]byte(body), &args)
	// log some trace
	log.Printf("COMMAND - INPUT - TYPE %v\nBODY: %v", model.GetType(), converted)
	return model.GetType(), model, args, nil
}

// Execute this command
func (p *Command) Execute(id string, body string) (string, error) {
	typ, command, args, _ := p.decode(id, body)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		return result.ToString(), nil
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		return result.ToString(), nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		return result.ToString(), nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		return result.ToString(), nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		return result.ToString(), nil
	default:
		log.Printf("Warning type %v is not implemented", typ)
	}
	return "", nil
}

// Test this command
func (p *Command) Test(id string, body string) (bool, error) {
	typ, command, args, _ := p.decode(id, body)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		return result.ToString() == "true", nil
		break
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		return result.ToString() == "true", nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		return result.ToString() == "true", nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		return result.ToString() == "true", nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		return result.ToString() == "true", nil
	default:
		log.Printf("Warning type %v is not implemented", typ)
	}
	return false, nil
}
