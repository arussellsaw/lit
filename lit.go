package lit

import (
	"encoding/json"
	"net/http"
	"text/template"
)

func LittleUI(wrapper, innerHTML string, fn func() (interface{}, error)) (http.Handler, error) {
	t := template.Must(template.New("lit").Parse(wrapper))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := page{
			InnerHTML: innerHTML,
		}
		v, err := fn()
		if err != nil {
			p.ErrString = err.Error()
			t.Execute(w, p)
			return
		}
		objectJSON, err := json.Marshal(v)
		p.ObjectJSON = string(objectJSON)
		if err != nil {
			p.ErrString = err.Error()
			t.Execute(w, p)
			return
		}
		err = t.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}), nil
}

type page struct {
	ObjectJSON string
	ErrString  string
	InnerHTML  string
}

var DefaultWrapper = `<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="https://cdnjs.com/libraries/bulma">
	<script src="https://cdn.jsdelivr.net/npm/vue"></script>
	<script type="text/javascript">
		var objectdata = {{.ObjectJSON}}
		var errString = "{{.ErrString}}"
	</script>
</head>
<body>
	{{.InnerHTML}}
	<script>
	var app = new Vue({
	  el: '#app',
	  data: {
	    objectData: objectdata,
	    error: errString,
	  }
	})
	</script>
</body>
</html>`
