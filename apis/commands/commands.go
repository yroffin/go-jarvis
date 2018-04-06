// Package commands for common apis
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
package commands

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/models"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	"github.com/yroffin/go-jarvis/apis/events"
	"github.com/yroffin/go-jarvis/services/chacon"
	"github.com/yroffin/go-jarvis/services/lua"
	"github.com/yroffin/go-jarvis/services/shell"
	"github.com/yroffin/go-jarvis/services/slack"
	"github.com/yroffin/go-jarvis/services/zway"
)

func init() {
	winter.Helper.Register("CommandBean", (&Command{}).New())
}

// Command internal members
type Command struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// Local cruds operations
	Crud interface{} `@crud:"/api/commands"`
	// Notification with injection mecanism
	LinkNotification events.INotification `@autowired:"NotificationBean" @link:"/api/commands" @href:"notifications"`
	Notification     events.INotification `@autowired:"NotificationBean"`
	// SlackService with injection mecanism
	SlackService slack.ISlackService `@autowired:"slack-service"`
	// ShellService with injection mecanism
	ShellService shell.IShellService `@autowired:"shell-service"`
	// LuaService with injection mecanism
	LuaService lua.ILuaService `@autowired:"lua-service"`
	// ChaconService with injection mecanism
	ChaconService chacon.IChaconService `@autowired:"chacon-service"`
	// ZwayService with injection mecanism
	ZwayService zway.IZwayService `@autowired:"zway-service"`
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
}

// ICommand implements IBean
type ICommand interface {
	engine.IAPI
}

// New constructor
func (p *Command) New() ICommand {
	bean := Command{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Command) Init() error {
	// Crud
	p.Factory = func() models.IPersistent {
		return (&CommandBean{}).New()
	}
	p.Factories = func() models.IPersistents {
		return (&CommandBeans{}).New()
	}
	p.HandlerTasksByID = func(id string, name string, body string) (interface{}, int, error) {
		if name == "execute" {
			// task
			return p.Execute(id, body)
		}
		if name == "test" {
			// task
			res, count, err := p.Test(id, body)
			return strconv.FormatBool(res), count, err
		}
		return "", -1, nil
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Command) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	return nil
}

// Validate this API
func (p *Command) Validate(name string) error {
	return nil
}

// Execute this command
func (p *Command) decode(id string, body string) (string, ICommandBean, map[string]interface{}, error) {
	// retrieve command and serialize it
	model := (&CommandBean{}).New()
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
func (p *Command) Execute(id string, body string) (interface{}, int, error) {
	typ, command, args, _ := p.decode(id, body)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		return result, -1, nil
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		return result, -1, nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		return result, -1, nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		return result, -1, nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		return result, -1, nil
	default:
		log.Printf("Warning type %v is not implemented", typ)
	}
	return "", -1, nil
}

// Test this command
func (p *Command) Test(id string, body string) (bool, int, error) {
	typ, command, args, _ := p.decode(id, body)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		return result.ToString() == "true", -1, nil
		break
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		return result.ToString() == "true", -1, nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		return result.ToString() == "true", -1, nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		return result.ToString() == "true", -1, nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		return result.ToString() == "true", -1, nil
	default:
		log.Printf("Warning type %v is not implemented", typ)
	}
	return false, -1, nil
}
