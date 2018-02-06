package main

import (
	"fmt"
	"net/http"

	"github.com/avct/lit"
)

func main() {
	litHandler, err := lit.LittleUI("example.html", func() (interface{}, error) {
		return foo{Foo: "foo", Bar: "bar"}, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	http.ListenAndServe(":8080", litHandler)
}

type foo struct {
	Foo string
	Bar string
}
