package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type ParameterParser interface {
	Parse(r io.Reader) (map[string]interface{}, error)
}

type JsonParser struct {}

func NewJsonParser() ParameterParser {
	return &JsonParser{}
}

type YamlParser struct {}

func NewYamlParser() ParameterParser {
	return &YamlParser{}
}

func (y *YamlParser) Parse(r io.Reader) (map[string]interface{}, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	if err := yaml.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (j *JsonParser) Parse(r io.Reader) (map[string]interface{}, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	if err := json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}

	return out, nil
}