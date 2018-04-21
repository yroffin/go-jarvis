// Package lua for common lua extends
// MIT License
//
// Copyright (c) 2018 yroffin
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
	"encoding/json"

	"github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// Loader module
func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"ObjectToJSON": objectToJSON,
	"ArrayToJSON":  arrayToJSON,
}

func objectToJSON(L *lua.LState) int {
	res := make(map[string]interface{})
	json.Unmarshal([]byte(L.CheckString(1)), &res)
	L.Push(luar.New(L, res))
	return 1
}

func arrayToJSON(L *lua.LState) int {
	res := make([]interface{}, 0)
	json.Unmarshal([]byte(L.CheckString(1)), &res)
	L.Push(luar.New(L, res))
	return 1
}
