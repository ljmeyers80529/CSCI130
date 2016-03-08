// web page homework #3
// larry meyere
// 08-Mar-2016

package main

import (
		"net/http"
		"io"
		"github.com/nu7hatch/gouuid"
		)

func main() {
	http.HandleFunc("/", root)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Session-Id")
	if err != nil {
		uuid, _ := uuid.NewV4()
		cookie = &http.Cookie {
			Name: "Session-Id",
			Value: uuid.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
		io.WriteString(res, "Session cookie created.")
	} else {
		io.WriteString(res, "Session cookie exists.")
		}
	io.WriteString(res, " - " + cookie.Value)
}
