// Larry Meyers
// 12-Mar-2016
//
// Web project, step #1

package main

import (
        "net/http"
        "html/template"
        )

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("htmlFiles/*.html"))
    }

func main() {
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/", root)
    http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
    tpl.Execute(res, nil)
}
