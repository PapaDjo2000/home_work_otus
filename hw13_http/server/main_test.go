package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	hendler(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	want := "привет"
	assert.Equal(t, want, recorder.Body.String())
}

func Test_get(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/get", nil)
	get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)

	contentType := recorder.Header().Get("Content-Type")
	assert.Equal(t, "application/json", contentType)

	var user User

	err := json.Unmarshal(recorder.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Ошибка десериализации JSON: %v", err)
	}

	want := User{Name: "dima", Age: 30}
	assert.Equal(t, want, user)
}

func TestPostHandler(t *testing.T) {
	testUser := User{Name: "maks", Age: 30}
	jsonData, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf("Ошибка сериализации JSON: %v", err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "", strings.TrimSpace(recorder.Body.String()))
}
