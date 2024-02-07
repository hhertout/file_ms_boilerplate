package service

import (
	json2 "encoding/json"
)

type Exporter struct{}

func NewExporter() Exporter {
	return Exporter{}
}

func (e Exporter) Json(json interface{}) ([]byte, error) {
	return json2.Marshal(json)
}

func (e Exporter) Csv() {

}

func (e Exporter) Xml() {

}
