// web page homework for https web site.
// larry meyere
// 12-Mar-2016

package main

import (
		"net/http"
                "log"
//		"io"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"

func main() {
	http.HandleFunc("/", root)
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
        if err != nil {
            log.Fatal(err)
        }
}

func root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	res.Header().Set(hdr, cont)
	res.Write([]byte("<H1>This is a sample https server.</H1>"))
        res.Write([]byte("<H2>This is served from local host</H2>"))
}
