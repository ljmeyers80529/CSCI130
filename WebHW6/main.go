// web page homework #3
// larry meyere
// 06-Mar-2016

package main

import (
		"net/http"
		"io"
		"strconv"
		)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}
	cookie, err := req.Cookie("Visit-Cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie {
			Name: "Visit-Cookie",
			Value: "0",
			}
		}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
