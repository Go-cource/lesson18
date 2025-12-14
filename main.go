package main

import (
	"net/http"
	_ "net/http/pprof"
)

const (
	maxSize = 1_000_000
)

func foo() {
	for {
		var s []int
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}

func main() {
	go foo()
	http.ListenAndServe(":8080", nil)
}
