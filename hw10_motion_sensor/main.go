package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func sensor(senschan chan<- int) {
	defer close(senschan)
	end := time.After(1*time.Minute + 1*time.Second)
	for {
		select {
		case <-end:
			return
		case <-time.After(1 * time.Second):
			limit := big.NewInt(50)
			n, err := rand.Int(rand.Reader, limit)
			if err != nil {
				continue
			}
			senschan <- int(n.Int64())
		}
	}
}

func data(senschan <-chan int, datachan chan<- float64) {
	defer close(datachan)
	data := make([]int, 0, 10)
	for value := range senschan {
		data = append(data, value)

		if len(data) == 10 {
			sum := 0

			for _, value := range data {
				sum += value
			}
			result := 0.0
			result = float64(sum) / float64(len(data))

			datachan <- result
			data = make([]int, 0, 10)
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
