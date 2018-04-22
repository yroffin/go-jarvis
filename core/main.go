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
package main

import (
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/winter"
	_ "github.com/yroffin/go-jarvis/core/apis/commands"
	_ "github.com/yroffin/go-jarvis/core/apis/configurations"
	_ "github.com/yroffin/go-jarvis/core/apis/connectors"
	_ "github.com/yroffin/go-jarvis/core/apis/datasources"
	_ "github.com/yroffin/go-jarvis/core/apis/devices"
	_ "github.com/yroffin/go-jarvis/core/apis/events"
	_ "github.com/yroffin/go-jarvis/core/apis/process"
	_ "github.com/yroffin/go-jarvis/core/apis/scripts"
	_ "github.com/yroffin/go-jarvis/core/apis/system"
	_ "github.com/yroffin/go-jarvis/core/apis/views"
	auto "github.com/yroffin/go-jarvis/core/auto"
	_ "github.com/yroffin/go-jarvis/core/services/chacon"
	_ "github.com/yroffin/go-jarvis/core/services/lua"
	_ "github.com/yroffin/go-jarvis/core/services/mqtt"
	_ "github.com/yroffin/go-jarvis/core/services/shell"
	_ "github.com/yroffin/go-jarvis/core/services/slack"
	_ "github.com/yroffin/go-jarvis/core/services/zway"
)

// PackInstance packer singleton
func PackInstance() winter.PackManager {
	auto.Pack = packr.NewBox("./dist")
	for _, res := range auto.Pack.List() {
		log.WithFields(log.Fields{
			"file": res,
		}).Info("Pack")
	}
	return auto.Pack
}

func main() {
	// Boot
	winter.Helper.Boot(PackInstance())
}
