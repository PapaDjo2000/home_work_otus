package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/yaml.v2"
)

type Book struct {
	ID     uuid.UUID `json:"ID" xml:"ID" yaml:"ID"`
	Title  string    `json:"Title" xml:"Title" yaml:"Title"`
	Author string    `json:"Author" xml:"Author" yaml:"Author"`
	Year   int       `json:"Year" xml:"Year" yaml:"Year"`
	Size   int       `json:"Size" xml:"Size" yaml:"Size"`
	Rate   float32   `json:"Rate" xml:"Rate" yaml:"Rate"`
	Sample json.RawMessage
}

func JsonMarshal(book []Book) ([]byte, error) {
	return json.Marshal(book)
}

func JsonUnMarshal(data []byte, book *[]Book) error {
	return json.Unmarshal(data, book)
}

func XmlMarshal(book []Book) ([]byte, error) {
	return xml.Marshal(book)
}

func XmlUnMarshal(data []byte, book *[]Book) error {
	return xml.Unmarshal(data, book)
}

func YamlMarshal(book []Book) ([]byte, error) {
	return yaml.Marshal(book)
}

func YamlUnMarshal(data []byte, book *[]Book) error {
	return yaml.Unmarshal(data, book)
}

func GobMarshal(book []Book) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(book)
	return buf.Bytes(), err
}

func GobUnmarshal(data []byte, book *[]Book) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(book)
	return err
}

func MsgPackMarshal(books []Book) ([]byte, error) {
	return msgpack.Marshal(books)
}

func MsgPackUnMarshal(data []byte, book *[]Book) error {
	return msgpack.Unmarshal(data, &book)
}

func main() {}
