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
	"strconv"

	log "github.com/sirupsen/logrus"

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
	// Execute this command
	Execute(id string, parameters map[string]interface{}) (interface{}, int, error)
	// Test this command
	Test(id string, parameters map[string]interface{}) (bool, int, error)
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
		var parameters = make(map[string]interface{})
		json.Unmarshal([]byte(body), &parameters)
		if name == "execute" {
			// task
			return p.Execute(id, parameters)
		}
		if name == "test" {
			// task
			res, count, err := p.Test(id, parameters)
			return strconv.FormatBool(res), count, err
		}
		return parameters, -1, nil
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
func (p *Command) decode(id string, parameters map[string]interface{}) (string, ICommandBean, map[string]interface{}, error) {
	// retrieve command and serialize it
	model, _ := p.GetByID(id)
	raw, _ := json.Marshal(&model)
	converted := make(map[string]interface{})
	json.Unmarshal(raw, &converted)
	// log some trace
	log.WithFields(log.Fields{
		"type": model.(ICommandBean).GetType(),
	}).Info("Execute command")
	log.WithFields(log.Fields{
		"model": model,
		"json":  models.ToJSON(converted),
	}).Debug("Execute command")
	return model.(ICommandBean).GetType(), model.(ICommandBean), parameters, nil
}

// Execute this command
func (p *Command) Execute(id string, parameters map[string]interface{}) (interface{}, int, error) {
	typ, command, args, _ := p.decode(id, parameters)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		log.WithFields(log.Fields{
			"json": models.ToJSON(result),
		}).Debug("Execute command - result")
		return result, -1, nil
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		log.WithFields(log.Fields{
			"json": models.ToJSON(result),
		}).Debug("Execute command - result")
		return result, -1, nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		log.WithFields(log.Fields{
			"json": models.ToJSON(result),
		}).Debug("Execute command - result")
		return result, -1, nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		log.WithFields(log.Fields{
			"json": models.ToJSON(result),
		}).Debug("Execute command - result")
		return result, -1, nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		log.WithFields(log.Fields{
			"json": models.ToJSON(result),
		}).Debug("Execute command - result")
		return result, -1, nil
	default:
		log.WithFields(log.Fields{
			"type": typ,
		}).Warn("Type not implemented")
	}
	return "", -1, nil
}

// Test this command
func (p *Command) Test(id string, parameters map[string]interface{}) (bool, int, error) {
	typ, command, args, _ := p.decode(id, parameters)
	switch typ {
	case "SLACK":
		result, _ := p.SlackService.AsObject(command, args)
		log.WithFields(log.Fields{
			"result": models.ToString(result),
		}).Info("Command test result")
		return models.ToString(result) == "true", -1, nil
		break
	case "SHELL":
		result, _ := p.ShellService.AsObject(command, args)
		log.WithFields(log.Fields{
			"result": models.ToString(result),
		}).Info("Command test result")
		return models.ToString(result) == "true", -1, nil
	case "LUA":
		result, _ := p.LuaService.AsObject(command, args)
		log.WithFields(log.Fields{
			"result": models.ToString(result),
		}).Info("Command test result")
		return models.ToString(result) == "true", -1, nil
	case "CHACON":
		result, _ := p.ChaconService.AsObject(command, args)
		log.WithFields(log.Fields{
			"result": models.ToString(result),
		}).Info("Command test result")
		return models.ToString(result) == "true", -1, nil
	case "ZWAY":
		result, _ := p.ZwayService.AsObject(command, args)
		log.WithFields(log.Fields{
			"result": models.ToString(result),
		}).Info("Command test result")
		return models.ToString(result) == "true", -1, nil
	default:
		log.WithFields(log.Fields{
			"type": typ,
		}).Warn("Type not implemented")
	}
	return false, -1, nil
}
