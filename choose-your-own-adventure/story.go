package adventure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

func init() {
	tmplt = template.Must(template.New("").Parse(defaultTemplate))
}

var tmplt *template.Template

//web page
var defaultTemplate = ` 
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>"choose your own adventure</title>
    </head>

    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
           <p>{{.}}</p>
        {{end}}

        <ul>
            {{range .Options}}
                <li>
                    <a href="/{{.Chapter}}">{{.Text}}</a>
                </li>
            {{end}}
        </ul>
    </body>
</html>`

// NewHandler ...
func NewHandler(s Story) http.Handler {
	return handler{}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tmplt.Execute(w, h.s["intro"])
	if err != nil {
		fmt.Println(err)
	}
}

// JSONStory ...
func JSONStory(r io.Reader) (Story, error) {
	dCoder := json.NewDecoder(r)
	var story Story
	if err := dCoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story ...
type Story map[string]Chapter

// Chapter ...
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option ...
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
