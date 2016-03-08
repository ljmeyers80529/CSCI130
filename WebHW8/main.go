// web page homework #8
// larry meyere
// 08-Mar-2016

package main

import (
		"net/http"
		"io"
		"fmt"
		"crypto/sha256"
		"crypto/hmac"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"
var body string = `<form method="POST">
					%s
					<input type="text" name="xname">
					<input type="Submit">
					</form>`
		
func main() {
	http.HandleFunc("/", root)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Session-Id")
	if err != nil {
		cookie = &http.Cookie {
			Name: "Session-Id",
			Value: "",
			//Secure: true,
			HttpOnly: true,
		}
	}
	if req.FormValue("xname") != "" {
		name := req.FormValue("xname")
		cookie.Value = name + "|" + hash(name)
		}
	http.SetCookie(res, cookie)
	res.Header().Set(hdr, cont)
	io.WriteString(res, fmt.Sprintf(body, cookie.Value))
}

func hash(data string) string {
	dHash := hmac.New(sha256.New, []byte("NotSoSecretKey"))
	io.WriteString(dHash, data)
	return fmt.Sprintf("%x", dHash.Sum(nil))
}