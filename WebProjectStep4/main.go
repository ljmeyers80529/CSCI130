// Larry Meyers
// 13-Mar-2016
//
// Web project, step #4

package main

import (
        "net/http"
        "html/template"
        "github.com/nu7hatch/gouuid"
        "log"
        "strings"
        "encoding/json"
        "encoding/base64"
        )

type user struct {
        Name string
        Age string
    }

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

// process web page root page.
func root(res http.ResponseWriter, req *http.Request) {
    sessionCookie, err := req.Cookie(cookieName)
    if err != nil {
        sessionCookie = cookieInitialValues()
    }
    setCookieInformation(req, sessionCookie)
    http.SetCookie(res, sessionCookie)
//    xs := strings.Split(sessionCookie.Value, "|")
    tpl.Execute(res, sessionCookie.Value)
}

// set cookie's initial uuid value.
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

// set or replace cookie information
func setCookieInformation(req *http.Request, cook *http.Cookie) {
    var uuid string = cook.Value
    
    if strings.Contains(cook.Value, "|") {
        uuid = strings.Split(cook.Value,"|")[0]
    }
    cook.Value = uuid + "|" + jsonEncodeBase64Data(req)
}

// save data from the form into a structured variable type, use JSON to marshal the information
// then encode that data into base 64.
func jsonEncodeBase64Data(req *http.Request) string {
    ui := user {
                Name: req.FormValue("username"),
                Age: req.FormValue("age"),
            }
    bs, err := json.Marshal(ui)
    if err != nil {
        log.Fatalln(err)
    }
    return base64.URLEncoding.EncodeToString(bs)
}

