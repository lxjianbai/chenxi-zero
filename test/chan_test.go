package test

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	var ch = make(chan int, 100)
	fmt.Println(111)
	<-ch
	fmt.Println(222)
}
