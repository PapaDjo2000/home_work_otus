package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSensor(t *testing.T) {
	senschan := make(chan int)
	go sensor(senschan)
	select {
	case num, ok := <-senschan:
		if !ok {
			t.Fatal()
		}
		assert.True(t, num >= 0 && num < 50)
	case <-time.After(1001 * time.Millisecond):
		t.Fatal()
	}
}

func TestData(t *testing.T) {
	sensch := make(chan int)
	datach := make(chan float64)

	go data(sensch, datach)

	var want float64
	for i := 0; i < 10; i++ {
		sensch <- i
		want += float64(i)
	}
	want /= 10
	result := <-datach

	assert.Equal(t, want, result)
}
