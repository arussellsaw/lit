package main

import (
	"fmt"
	"net/http"

	"github.com/avct/lit"
	"github.com/gobuffalo/packr"
)

func main() {
	b := packr.NewBox("./static")
	litHandler, err := lit.LittleUI(b.String("example.html"), func(r *http.Request) (interface{}, error) {
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
