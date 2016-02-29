// web page homework #1
// larry meyere
// 29-Feb-2016

package main

import (
		"net/http"
		"io"
		"strings"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path,"/")
	res.Header().Set(hdr, cont)
	io.WriteString(res, "<h1>URL path = /"+fs[1]+"</h1>")
}
