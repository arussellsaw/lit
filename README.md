# 🔥 lit 🔥

#### Little Ui Toolkit

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

	"github.com/avct/lit"
	"github.com/gobuffalo/packr"
)

func main() {
	b := packr.newbox("./static")
	lithandler, err := lit.littleui(b.string("example.html"), func() (interface{}, error) {
		return foo{foo: "foo", bar: "bar"}, nil
	})
	if err != nil {
		fmt.println(err)
		return
	}
	http.listenandserve(":8080", lithandler)
}

type foo struct {
	foo string
	bar string
}
```

html:
```html
<div id="app" class="button is-primary">
  {{ objectdata.foo }}
</div>
```
