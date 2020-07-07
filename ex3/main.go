package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	story, err := jsonParse(file)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("").Parse(webpage)
	if err != nil {
		panic(err)
	}
	h := handler{
		a: story,
		t: tmpl,
	}
	http.ListenAndServe(":8080", h)
}

const webpage = `
<title>{{.Title}}</title>
<body>
<h1>{{.Title}}</h1>
{{range .Story}}
<p>{{.}}</p>
{{end}}
<ul>
{{range .Options}}
<li><a href="/{{.Arc}}">{{.Text}}</a></li>
{{end}}
</ul>
</body>
`

type handler struct {
	a Adventure
	t *template.Template
}

// uses the the template defined in webpage to generate generate html code
// based on the key stored in the Adventure field a of the handler struct
// then sends the html to the browser
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/")
	h.t.Execute(w, h.a[path])
}

type Adventure map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func jsonParse(r io.Reader) (Adventure, error) {
	d := json.NewDecoder(r)
	story := make(Adventure)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
