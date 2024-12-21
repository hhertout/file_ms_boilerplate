package service

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
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

func (e Exporter) Csv(entries []map[string]interface{}) (bytes.Buffer, error) {
	var headers []string
	for _, e := range entries {
		for k := range e {
			if !arrayContain(k, headers) {
				headers = append(headers, k)
			}
		}
	}

	var buffer bytes.Buffer
	w := csv.NewWriter(&buffer)
	if err := w.Write(headers); err != nil {
		return bytes.Buffer{}, err
	}

	for _, e := range entries {
		values := make([]string, len(headers))
		for k, v := range e {
			index, err := IndexOf(k, headers)
			if err != nil {
				return bytes.Buffer{}, err
			}
			values[index] = toString(v)
		}

		if err := w.Write(values); err != nil {
			return bytes.Buffer{}, err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return bytes.Buffer{}, err
	}

	return buffer, nil
}

func arrayContain[T comparable](search T, slice []T) bool {
	for _, e := range slice {
		if e == search {
			return true
		}
	}
	return false
}

func toString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", value)
}

func IndexOf(search string, slice []string) (int, error) {
	for i, v := range slice {
		if search == v {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}
