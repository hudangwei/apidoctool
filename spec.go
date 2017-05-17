package main

import (
	"encoding/json"
	"strings"
)

func DecodeJSON(data []byte) (*API, error) {
	var api API

	s := strings.Replace(string(data), "$ref", "myRef", -1)
	err := json.Unmarshal([]byte(s), &api)

	definitions := make(map[string]Definition)

	for k, v := range api.Definitions {
		definitions["#/definitions/"+k] = v
	}
	api.Definitions = definitions
	return &api, err
}

/**
 * swagger全局配置
 */
type Info struct {
	Version        string  `json,yaml:"version"`
	Title          string  `json,yaml:"title"`
	Description    string  `json,yaml:"description"`
	TermsOfService string  `json,yaml:"termsOfService"`
	Contact        Contact `json,yaml:"contact"`
	License        License `json,yaml:"license"`
}

/**
 * 联系方式
 */
type Contact struct {
	Name  string `json,yaml:"name"`
	Email string `json,yaml:"email"`
	URL   string `json,yaml:"url"`
}

/**
 * 协议
 */
type License struct {
	Name string `json,yaml:"name"`
	URL  string `json,yaml:"url"`
}

type Tag struct {
	Name        string `json,yaml:"name"`
	Description string `json,yaml:"description"`
}

type API struct {
	Swagger     string                `json,yaml:"swagger"` //swagger版本
	Info        Info                  `json,yaml:"info"`
	Host        string                `json,yaml:"host"`
	BasePath    string                `json,yaml:"basePath"`
	Tags        []Tag                 `json,yaml:"tags"`
	Paths       map[string]Methods    `json,yaml:"paths"`
	Definitions map[string]Definition `json,yaml:"definitions"`
}

type Methods struct {
	Get    *Operation `json,yaml:"get"`
	Post   *Operation `json,yaml:"post"`
	Put    *Operation `json,yaml:"put"`
	Patch  *Operation `json,yaml:"patch"`
	Delete *Operation `json,yaml:"delete"`
	Head   *Operation `json,yaml:"head"`
}

type Operation struct {
	Tags        []string            `json,yaml:"tags"`
	Summary     string              `json,yaml:"summary"`
	Description string              `json,yaml:"description"`
	OperationID string              `json,yaml:"operationId"`
	Consumes    []string            `json,yaml:"consumes"`
	Produces    []string            `json,yaml:"produces"`
	Parameters  []Parameter         `json,yaml:"parameters"`
	Responses   map[string]Response `json,yaml:"responses"`
}

type Parameter struct {
	Name        string            `json,yaml:"name"`
	In          string            `json,yaml:"in"`
	Description string            `json,yaml:"description"`
	Required    bool              `json,yaml:"required"`
	Type        string            `json,yaml:"type"`
	Format      string            `json,yaml:"format"`
	Schema      map[string]string `json,yaml:"schema"`
}

type Response struct {
	Type        string                 `json,yaml:"type"`
	Description string                 `json,yaml:"description"`
	Schema      map[string]interface{} `json,yaml:"schema"`
}

type Definition struct {
	Title      string               `json,yaml:"title"`
	Type       string               `json,yaml:"type"`
	Properties map[string]Propertie `json,yaml:"properties"`
}

type Propertie struct {
	Description string            `json,yaml:"description"`
	Type        string            `json,yaml:"type"`
	Format      string            `json,yaml:"format"`
	MyRef       string            `json,yaml:"myRef"`
	Items       map[string]string `json,yaml:"items"`
}
