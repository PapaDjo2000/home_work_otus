package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func sensor(senschan chan<- int) {
	defer close(senschan)
	end := time.After(1 * time.Minute)
	limit := big.NewInt(50)

	for {
		n, err := rand.Int(rand.Reader, limit)
		if err != nil {
			continue
		}
		select {
		case <-end:
			return
		case senschan <- int(n.Int64()):
		}
	}
}

func data(senschan <-chan int, datachan chan<- float64) {
	defer close(datachan)
	var data float64
	var count int
	for value := range senschan {
		data += float64(value)
		count++
		if count == 10 {
			datachan <- data / 10
			data = 0
			count = 0
		}
	}
}

func main() {
	senschan := make(chan int)
	datachan := make(chan float64)

	go sensor(senschan)
	go data(senschan, datachan)

	for value := range datachan {
		fmt.Println(value)
	}
}
