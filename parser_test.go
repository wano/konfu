package main

import (
	"os"
	"testing"
)

func TestYamlParser_Parse(t *testing.T) {
	parser := NewYamlParser()
	f, err := os.Open("test/sample.yml")
	if err != nil {
		t.Fatal(err)
	}

	got, err := parser.Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	name, ok := got["name"]
	if !ok {
		t.Errorf("name key not found")
	}

	if name != "takuya" {
		t.Errorf("got: takuya, got: %s", name)
	}
}

func TestJsonParser_Parse(t *testing.T) {
	parser := NewJsonParser()
	f, err := os.Open("test/sample.json")
	if err != nil {
		t.Fatal(err)
	}

	got, err := parser.Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	name, ok := got["name"]
	if !ok {
		t.Errorf("name key not found")
	}

	if name != "takuya" {
		t.Errorf("got: takuya, got: %s", name)
	}
}