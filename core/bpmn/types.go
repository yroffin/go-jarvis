// Package bpmn module
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
package bpmn

import "encoding/xml"

type bpmnXML struct {
	XMLName xml.Name       `xml:"definitions"`
	Process bpmnProcessXML `xml:"process"`
}

type bpmnProcessXML struct {
	XMLName xml.Name             `xml:"process"`
	ID      string               `xml:"id,attr"`
	Start   []bpmnStartEventXML  `xml:"startEvent"`
	End     []bpmnEndEventXML    `xml:"endEvent"`
	Flow    []bpmnFlowXML        `xml:"sequenceFlow"`
	Service []bpmnServiceTaskXML `xml:"serviceTask"`
}

type bpmnStartEventXML struct {
	XMLName xml.Name          `xml:"startEvent"`
	ID      string            `xml:"id,attr"`
	Name    string            `xml:"name,attr"`
	Out     []bpmnOutgoingXML `xml:"outgoing"`
}

type bpmnOutgoingXML struct {
	XMLName xml.Name `xml:"outgoing"`
	Body    string   `xml:",chardata"`
}

type bpmnEndEventXML struct {
	XMLName xml.Name          `xml:"endEvent"`
	ID      string            `xml:"id,attr"`
	Name    string            `xml:"name,attr"`
	Inc     []bpmnIncomingXML `xml:"incoming"`
}

type bpmnIncomingXML struct {
	XMLName xml.Name `xml:"incoming"`
	Body    string   `xml:",chardata"`
}

type bpmnFlowXML struct {
	XMLName xml.Name `xml:"sequenceFlow"`
	ID      string   `xml:"id,attr"`
	Source  string   `xml:"sourceRef,attr"`
	Target  string   `xml:"targetRef,attr"`
}

type bpmnServiceTaskXML struct {
	XMLName   xml.Name           `xml:"serviceTask"`
	ID        string             `xml:"id,attr"`
	Extention []bpmnExtentionXML `xml:"extensionElements"`
}

type bpmnExtentionXML struct {
	XMLName        xml.Name                `xml:"extensionElements"`
	ID             string                  `xml:"id,attr"`
	TaskDefinition []bpmnTaskDefinitionXML `xml:"taskDefinition"`
}

type bpmnTaskDefinitionXML struct {
	XMLName xml.Name `xml:"taskDefinition"`
	Type    string   `xml:"type,attr"`
}

// Event internal event
type Event struct {
	Process *ProcessInstance
	Node    *NodeInstance
}

// EventBpmn internal event
type EventBpmn struct {
	ID  string `json:"id"`
	Out string `json:"outgoing,omitempty"`
	Inc string `json:"incoming,omitempty"`
}

// ProcessBpmn processus
type ProcessBpmn struct {
	ID     string               `json:"id"`
	Tasks  map[string]TaskBpmn  `json:"tasks"`
	Events map[string]EventBpmn `json:"events"`
	Flows  map[string]FlowBpmn  `json:"flows"`
}

// TaskBpmn internal task
type TaskBpmn struct {
	ID string `json:"id"`
}

// FlowBpmn internal flow
type FlowBpmn struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// ProcessInstance processus
type ProcessInstance struct {
	ID         string         `json:"id"`
	Active     bool           `json:"active"`
	Compiled   bool           `json:"compiled"`
	Definition ProcessBpmn    `json:"definition"`
	Nodes      []NodeInstance `json:"nodes"`
	Edges      []EdgeInstance `json:"edges"`
	Current    string         `json:"current"`
}

// NodeInstance internal task
type NodeInstance struct {
	ID          string `json:"id"`
	Ref         string `json:"reference"`
	IsActive    bool   `json:"isActive"`
	IsEvent     bool   `json:"isEvent"`
	IsTask      bool   `json:"isTask"`
	IsStart     bool   `json:"isStart"`
	IsCompleted bool   `json:"isCompleted"`
}

// EdgeInstance internal task
type EdgeInstance struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}
