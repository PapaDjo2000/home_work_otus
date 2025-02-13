package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var user User

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func httpGet(url string, user *User) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}
	if contentType == "application/json" {
		err = json.Unmarshal(body, user)
		if err != nil {
			fmt.Println("Ошибка десериализации", err)
		}
		fmt.Println("получен", user)
		return
	} else {
		fmt.Printf("получен %s.\n", string(body))
		return
	}
}

func httpPost(url string, userserv User) {
	body, err := json.Marshal(userserv)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}
	resp, err := http.Post(
		url,
		"application/json",
		strings.NewReader(string(body)),
	)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}
}
func main() {
	var urlServ, endPoint string
	userserv := User{"maks", 32}

	flag.StringVar(&urlServ, "urlServ", "http://localhost:8080", "url address")
	flag.StringVar(&endPoint, "endPoint", "", "next restful andpoint")
	flag.Parse()

	url := urlServ + endPoint
	httpGet(url, &user)
	httpPost(url, userserv)
}
