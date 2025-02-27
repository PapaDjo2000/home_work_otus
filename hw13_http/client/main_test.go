package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_HttpGet(t *testing.T) {
	testUser := User{Name: "maks", Age: 30}
	jsonData, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf("ошибка при создании JSON: %v", err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}))
	defer ts.Close()

	var usertest User
	httpGet(ts.URL, &usertest)

	assert.Equal(t, testUser, usertest)
}

func TestClient_HttpPost(t *testing.T) {
	testUser := User{Name: "maks", Age: 30}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var newUser User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding JSON: %v", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()
	httpPost(ts.URL, testUser)
	t.Logf("Test passed: Data sent successfully to server.")
}
