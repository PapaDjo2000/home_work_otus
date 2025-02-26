package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var user User

func HttpGet(urlserv string, user *User) {
	urlserv1, err := url.Parse(urlserv)
	if err != nil {
		fmt.Print("неверный урл адрес", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", urlserv1.String(), nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
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
	}
	fmt.Printf("получен %s\n", string(body))
}

func HttpPost(urlserv string, userserv User) {
	urlserv1, err := url.Parse(urlserv)
	if err != nil {
		fmt.Print("неверный урл адрес", err)
		return
	}
	body, err := json.Marshal(userserv)
	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", urlserv1.String(), strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}
	fmt.Println("отправлен на сервер ", string(body))
}

func main() {
	var urlServ, endPoint string
	userserv := User{"maks", 32}

	flag.StringVar(&urlServ, "urlServ", "http://localhost:8080", "urlServ address")
	flag.StringVar(&endPoint, "endPoint", "", "next restful andPoint")
	flag.Parse()

	url1 := urlServ + endPoint

	switch url1 {
	case "http://localhost:8080/get":
		HttpGet(url1, &user)
	case "http://localhost:8080/post":
		HttpPost(url1, userserv)
	default:
		HttpGet(url1, &user)
	}
}
