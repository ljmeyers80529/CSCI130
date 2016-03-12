// Larry Meyers
// 12-Mar-2016
//
// Web project, step #2

package main

import (
        "net/http"
        "html/template"
	"github.com/nu7hatch/gouuid"
        "log"
        )

var tpl *template.Template
var cookieName = "session-fino"
var sessionCookie *http.Cookie

func init() {
    tpl = template.Must(template.ParseGlob("htmlFiles/*.html"))
    }

func main() {
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/", root)
    http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
    sessionCookie, err := req.Cookie(cookieName)
    if err != nil {
        sessionCookie = cookieInitialValues()
    }
    http.SetCookie(res, sessionCookie)
    tpl.Execute(res, nil)
}
   
func cookieInitialValues() *http.Cookie {
    uuid, err := uuid.NewV4()
    if err != nil {
        log.Fatalln(err)
    }
    cookie := &http.Cookie {
            Name: cookieName,
            Value: uuid.String(),
            //Secure: true,
            HttpOnly: true,
            }
    return cookie
}