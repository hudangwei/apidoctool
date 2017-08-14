package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GenerateMarkDownFile(filename string) error {
	var b []byte
	var err error
	if strings.Contains(filename, "http://") {
		b, err = ReadFromURL(filename)
	} else {
		b, err = ReadFromJson(filename)
	}

	if err != nil {
		return err
	}

	api, err := DecodeJSON(b)
	if err != nil {
		return err
	}

	currpath, _ := os.Getwd()
	for _, v := range api.Tags {
		mdPath := path.Join(currpath, v.Name)
		os.MkdirAll(mdPath, 0755)
	}
	apiservice := NewAPIService(api)

	weightMap := make(map[string]int)
	for _, serviceInterface := range apiservice.ServiceInterfaces {
		serviceInterface.OperationID = strings.Replace(serviceInterface.OperationID, ".", "_", -1)
		weightMap[serviceInterface.Tag]++
		markdownDoc := ""
		hugoHeader := fmt.Sprintf("+++\ntitle=\"%s\"\nurl=\"/%s/%s\"\n[menu.side]\n  name=\"%s\"\n  parent=\"%s\"\n  weight=%d\n+++\n\n",
			serviceInterface.Description, serviceInterface.Tag, strings.ToLower(serviceInterface.OperationID), serviceInterface.Description, serviceInterface.Tag, weightMap[serviceInterface.Tag])
		markdownDoc += hugoHeader
		desc := fmt.Sprintf("## %s\n\n", serviceInterface.Description)
		markdownDoc += desc
		method := fmt.Sprintf("**请求方式**\n\n`%s`\n\n", serviceInterface.Method)
		markdownDoc += method
		uri := fmt.Sprintf("**请求路径**\n\n`%s`\n\n", api.BasePath+serviceInterface.URI)
		markdownDoc += uri
		if len(serviceInterface.Params) > 0 {
			markdownDoc += "**请求参数说明**\n\n|参数名|必选|参数方式|类型|说明|\n| :---- | :--- | :----- | :----- | ---------------------------------------- |\n"
			for _, parameter := range serviceInterface.Params {
				param := fmt.Sprintf("|%s|%t|%s|%s|%s|\n", parameter.Name, parameter.Required, parameter.In, parameter.Type, parameter.Description)
				markdownDoc += param
			}
		}

		if len(serviceInterface.Body) > 0 {
			markdownDoc += "\n**返回参数说明**\n\n|参数名|类型|说明|\n| :----- | :----- | ------------------------ |\n"
			for _, v := range serviceInterface.Body {
				param := fmt.Sprintf("|%s|%s|%s|\n", v.Name, v.Type, v.Description)
				markdownDoc += param
			}
		}

		markdownDoc += parseSchema(api, serviceInterface.Definitions)

		fpath := path.Join(serviceInterface.Tag, strings.ToLower(serviceInterface.OperationID)+".md")

		ioutil.WriteFile(fpath, []byte(markdownDoc), 0666)
	}

	return nil
}

func parseSchema(api *API, defs map[string]Definition) string {
	markdownDoc := ""
	obj := make(map[string]Definition)
	for k, v := range defs {
		if len(v.Properties) > 0 {
			markdownDoc += fmt.Sprintf("\n**%s 数据结构**\n\n|成员|类型|说明|\n| :----- | :----- | ------------------------ |\n", k)
			for k, v := range v.Properties {
				if v.MyRef != "" {
					v.Type = api.Definitions[v.MyRef].Type + " (" + api.Definitions[v.MyRef].Title + ")"
					obj[api.Definitions[v.MyRef].Title] = api.Definitions[v.MyRef]

					if _, ok := defs[api.Definitions[v.MyRef].Title]; ok {
						delete(obj, api.Definitions[v.MyRef].Title)
					}
				}
				if v.Type == "array" {
					if _, ok := v.Items["myRef"]; ok {
						v.Type = "array (" + api.Definitions[v.Items["myRef"]].Title + ")"
						obj[api.Definitions[v.Items["myRef"]].Title] = api.Definitions[v.Items["myRef"]]

						if _, ok := defs[api.Definitions[v.Items["myRef"]].Title]; ok {
							delete(obj, api.Definitions[v.Items["myRef"]].Title)
						}
					}
				}
				param := fmt.Sprintf("|%s|%s|%s|\n", k, v.Type, v.Description)
				markdownDoc += param
			}
		}
	}
	if len(obj) == 0 {
		return markdownDoc
	}
	return markdownDoc + parseSchema(api, obj)
}
