package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/yaml.v2"
)

var books = []Book{{
	ID:     uuid.New(),
	Title:  "Властелин колец 2",
	Author: "Толкин",
	Year:   1955,
	Size:   954,
	Rate:   154.32,
}}

func TestSerialisationDeserialisation(t *testing.T) {
	var desBook []Book

	//Test Json

	want, err := json.Marshal(books)
	assert.NoError(t, err)
	result, err := JSONMarshal(books)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = json.Unmarshal(result, &desBook)
	assert.NoError(t, err)
	err = JSONUnMarshal(result, &desBook)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	//Test Xml

	want, err = xml.Marshal(books)
	assert.NoError(t, err)
	result, err = XMLMarshal(books)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = xml.Unmarshal(result, &desBook)
	assert.NoError(t, err)
	err = XMLUnMarshal(result, &desBook)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	//Test Yaml

	want, err = yaml.Marshal(books)
	assert.NoError(t, err)
	result, err = YAMLMarshal(books)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = yaml.Unmarshal(result, &desBook)
	assert.NoError(t, err)
	err = YAMLUnMarshal(result, &desBook)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	//Test Msgpack

	want, err = msgpack.Marshal(books)
	assert.NoError(t, err)
	result, err = MsgPackMarshal(books)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = msgpack.Unmarshal(result, &desBook)
	assert.NoError(t, err)
	err = MsgPackUnMarshal(result, &desBook)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	//Test GOB

	var buf bytes.Buffer
	//Encoder
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(books)
	assert.NoError(t, err)
	want = buf.Bytes()

	result, err = GobMarshal(books)
	assert.NoError(t, err)
	assert.Equal(t, want, result)
	//Decoder
	newbuf := bytes.NewBuffer(want)
	decod := gob.NewDecoder(newbuf)
	err = decod.Decode(&desBook)
	assert.NoError(t, err)

	err = GobUnmarshal(result, &desBook)
	assert.NoError(t, err)
	assert.Equal(t, books, desBook)
}
