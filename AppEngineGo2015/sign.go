package hello

import (
	"fmt"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/sign", sign)
}

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content")) // HL
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML)) // HL

const signTemplateHTML = `
<html><body><p>You wrote: {{.}}</p></body></html>
`
