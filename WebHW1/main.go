// web page homework #1
// larry meyere
// 29-Feb-2016

package main

import (
		"net/http"
		"io"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":9000", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	res.Header().Set(hdr, cont)
	io.WriteString(res, "<h1>URL path = "+req.URL.Path+"</h1>")
}
