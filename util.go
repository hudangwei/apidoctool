package main

import (
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/httplib"
)

func ReadFromJson(file_name string) ([]byte, error) {
	fin, err := os.Open(file_name)
	defer fin.Close()
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(fin)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func ReadFromURL(url string) ([]byte, error) {
	req := httplib.Get(url)
	return req.Bytes()
}
