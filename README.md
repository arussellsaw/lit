# 🔥 lit 🔥

#### Little Interface Toolkit

Lit is a toolkit to make your life easier when adding small UI endpoints to your services.
Lit works via two components, a `Lit Wrapper` responsible for constructing the html, and injecting
data into the page via templating, and `Inner HTML` which handles rendering the data injected by
the wrapper. The current `DefaultWrapper` uses VueJS and Bulma.

#### Example:

Go:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/arussellsaw/lit"
	"github.com/gobuffalo/packr"
)

func main() {
	b := packr.NewBox("./static")
	litHandler, err := lit.LittleUI(lit.DefaultWrapper, b.String("example.html"), func(r *http.Request) (interface{}, error) {
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
```

html:
```html
<div id="app" class="button is-primary">
  {{ objectdata.foo }}
</div>
```
