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
	"github.com/yroffin/go-boot-sqllite/core/winter"
	_ "github.com/yroffin/go-jarvis/apis/commands"
	_ "github.com/yroffin/go-jarvis/apis/configurations"
	_ "github.com/yroffin/go-jarvis/apis/connectors"
	_ "github.com/yroffin/go-jarvis/apis/datasources"
	_ "github.com/yroffin/go-jarvis/apis/devices"
	_ "github.com/yroffin/go-jarvis/apis/events"
	_ "github.com/yroffin/go-jarvis/apis/process"
	_ "github.com/yroffin/go-jarvis/apis/scripts"
	_ "github.com/yroffin/go-jarvis/apis/system"
	_ "github.com/yroffin/go-jarvis/apis/views"
	_ "github.com/yroffin/go-jarvis/auto"
	_ "github.com/yroffin/go-jarvis/services/chacon"
	_ "github.com/yroffin/go-jarvis/services/lua"
	_ "github.com/yroffin/go-jarvis/services/mqtt"
	_ "github.com/yroffin/go-jarvis/services/shell"
	_ "github.com/yroffin/go-jarvis/services/slack"
	_ "github.com/yroffin/go-jarvis/services/zway"
)

func packInstance() winter.PackManager {
	box := packr.NewBox("./dist")
	return box
}

func main() {
	// Boot
	winter.Helper.Boot(packInstance())
}
