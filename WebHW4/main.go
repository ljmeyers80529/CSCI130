// web page homework #3
// larry meyere
// 06-Mar-2016

package main

import (
		"net/http"
		"io"
		"fmt"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"
var body string = `<form method="GET">
					<input type="text" name="n">
					<input type="Submit">
					</form>`

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}
	res.Header().Set(hdr, cont)
	key := "n"
	value := req.FormValue(key)
	fmt.Println("Returned value = "+value)
	io.WriteString(res, body)
	if (req.Method == "GET") && (value != "") {
		io.WriteString(res, "<h1>Value entered = " + value + "</h1>")
	}
}
