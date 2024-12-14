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
	ID     uuid.UUID `json:"id" xml:"id" yaml:"id" bson:"id"`
	Title  string    `json:"title" xml:"title" yaml:"title" bson:"title"`
	Author string    `json:"author" xml:"author" yaml:"author" bson:"author"`
	Year   int       `json:"year" xml:"year" yaml:"year" bson:"year"`
	Size   int       `json:"size" xml:"size" yaml:"size" bson:"size"`
	Rate   float32   `json:"rate" xml:"rate" yaml:"rate" bson:"rate"`
	Sample json.RawMessage
}

func JSONMarshal(book []Book) ([]byte, error) {
	return json.Marshal(book)
}

func JSONUnMarshal(data []byte, book *[]Book) error {
	return json.Unmarshal(data, book)
}

func XMLMarshal(book []Book) ([]byte, error) {
	return xml.Marshal(book)
}

func XMLUnMarshal(data []byte, book *[]Book) error {
	return xml.Unmarshal(data, book)
}

func YAMLMarshal(book []Book) ([]byte, error) {
	return yaml.Marshal(book)
}

func YAMLUnMarshal(data []byte, book *[]Book) error {
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
