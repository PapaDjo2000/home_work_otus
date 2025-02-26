package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func hendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "привет")
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.URL)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.URL)

	w.Header().Set("Content-Type", "application/json")
	user := User{
		Name: "dima",
		Age:  30,
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("ошибка записи Json", err)
	}
	fmt.Println("отправлен ", user)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.URL)

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Получен ", newUser)
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()
	http.HandleFunc("/", hendler)
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
	server := &http.Server{
		Addr:         *ip + ":" + *port,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on", *ip+":"+*port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
