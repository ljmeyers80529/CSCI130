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
                "strings"
		)

var hdr string = "Content-Type"
var cont string = "text/html; charset=utf-8"
var body string = `<form method="POST">
					%s
					<input type="text" name="xname">
					<input type="Submit">
					</form>
					%s`
		
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
	io.WriteString(res, fmt.Sprintf(body, cookie.Value, checkCookie(res, req.FormValue("xname"), cookie.Value)))
}

func hash(data string) string {
	dHash := hmac.New(sha256.New, []byte("NotSoSecretKey"))
	io.WriteString(dHash, data)
	return fmt.Sprintf("%x", dHash.Sum(nil))
}

func checkCookie(res http.ResponseWriter, frmVal, coValue string) string {
    if coValue == "" {
        return ``
    }
    xs := strings.Split(coValue, "|")
    xs[0] = xs[0] + "tampered"
    var v = "Valid"
    if xs[1] != hash(xs[0]) {
        v = "InValid"
    }
    return fmt.Sprintf("<h1> Cookie values<br>index 1 = %s<br>index 2 = %s<br>Cookie is %s", xs[0], xs[1], v)
}