package lit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

func LittleUI(htmlPath string, fn func() (interface{}, error)) (http.Handler, error) {
	t := template.Must(template.New("lit").Parse(tmplString))
	f, err := os.Open(htmlPath)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := page{
			InnerHTML: string(buf),
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

var tmplString = `<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="http://cdn.staging.avocet.io/internal-assets/styles.css">
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
