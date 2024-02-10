package tests

import (
	json2 "encoding/json"
	"github.com/eco-challenge/src/service"
	"testing"
)

type TestJson struct {
	Foo string `json:"foo" xml:"foo"`
}

func TestExporterJson(t *testing.T) {
	t.Run("json test", func(t *testing.T) {
		exporter := service.NewExporter()
		jsonData := TestJson{
			Foo: "bar",
		}
		decodedData, err := exporter.Json(jsonData)
		if err != nil {
			t.Errorf("Failed to parse to json: %v", err)
		}
		expected, err := json2.Marshal(jsonData)
		if err != nil {
			t.Errorf("Failed to parse expected result: %v", err)
		}

		if string(decodedData) != string(expected) {
			t.Errorf("Data don't match. Expect : %s, Result : %s", expected, string(decodedData))
		}
	})

	t.Run("xml test", func(t *testing.T) {
		exporter := service.NewExporter()
		jsonData := TestJson{
			Foo: "bar",
		}
		decodedData, err := exporter.JsonToXml(jsonData)
		if err != nil {
			t.Errorf("Failed to parse to json: %v", err)
		}
		expected := "<foo>bar</foo>"

		if decodedData != expected {
			t.Errorf("Data don't match. Expect : %s, Result : %s", expected, decodedData)
		}
	})
}
