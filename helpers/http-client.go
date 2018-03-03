// Package models for all models
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
package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPClient simple command model
type HTTPClient struct {
	// URL
	URL string `json:"url"`
	// User
	User string `json:"user"`
	// Password
	Password string `json:"password"`
}

// Call http
func (p *HTTPClient) request(method string, path string, body map[string]interface{}) (*http.Request, error) {
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			log.Println("Body/Error:", err)
			return nil, err
		}
		return http.NewRequest(method, p.URL+path, strings.NewReader(string(payload)))
	}
	return http.NewRequest(method, p.URL+path, nil)
}

// Call http
func (p *HTTPClient) Call(method string, path string, body map[string]interface{}, headers map[string]string, params map[string]string) (map[string]interface{}, error) {
	// build client
	req, err := p.request(method, path, body)

	if err != nil {
		log.Println("NewRequest/Error:", err)
		return nil, err
	}
	client := &http.Client{}

	// fix headers
	//req.Header.Add("Authorization", "secretToken")

	// execute request
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Do/Error:", err)
		return nil, err
	}

	// read stream
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Read/Error:", err)
		return nil, err
	}

	log.Println("Url:", req.URL, "Body:", body, "Status:", resp.Status)

	// build result
	args := make(map[string]interface{})
	json.Unmarshal(data, &args)
	return args, nil
}

// Call http
func (p *HTTPClient) GET(path string, headers map[string]string, params map[string]string) (map[string]interface{}, error) {
	return p.Call("GET", path, nil, headers, params)
}

// Call http
func (p *HTTPClient) POST(path string, body map[string]interface{}, headers map[string]string, params map[string]string) (map[string]interface{}, error) {
	return p.Call("POST", path, body, headers, params)
}
