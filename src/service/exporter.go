package service

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Exporter struct{}

func NewExporter() Exporter {
	return Exporter{}
}

func (e Exporter) Json(input any) ([]byte, error) {
	return json.Marshal(input)
}

func (e Exporter) JsonToXml(input interface{}) (string, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return "", err
	}

	return e.parseToXml(result)
}

func (e Exporter) parseToXml(entry map[string]interface{}) (string, error) {
	formattedXml := ""
	for k, v := range entry {
		formattedXml += fmt.Sprintf("<%v>", k)
		if reflect.TypeOf(v).String() != "string" {
			var parse map[string]interface{}
			jsonData, err := json.Marshal(v)
			if err != nil {
				return "", err
			}

			if err := json.Unmarshal(jsonData, &parse); err != nil {
				return "", err
			}

			toXml, err := e.parseToXml(parse)
			if err != nil {
				return "", err
			}

			formattedXml += toXml
		} else if reflect.TypeOf(v).String() == "string" {
			s := fmt.Sprintf("%v", v)
			formattedXml += s
		}
		formattedXml += fmt.Sprintf("</%v>", k)
	}
	return formattedXml, nil
}

func (e Exporter) Csv() {

}
