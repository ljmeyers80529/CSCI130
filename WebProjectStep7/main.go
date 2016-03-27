package main

import (
    "net/http"
    "html/template"
    "strings"
    )

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("htmlFiles/*.html"))
}

func main() {
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/", rootPage)
    http.HandleFunc("/login", loginPage)
    http.ListenAndServe(":8080", nil)
}

// process root / top page.
func rootPage(res http.ResponseWriter, req *http.Request) {
    sessionCookie := retrieveCookie(res, req)
    if req.Method == "POST" {
        updateCookie(req, sessionCookie)
    }
    http.SetCookie(res, sessionCookie)
    
    var cookieState = "False"
    if isCookieInformationValid(sessionCookie) {
        cookieState = "True"
    }
    cookieInfo := cookieInformation {
        State: cookieState,
        OriginalCookieValue: sessionCookie.Value,
        OriginalData: decodeUserInformation(strings.Split(sessionCookie.Value, "|")[1]),
    }
    tpl.ExecuteTemplate(res, "root.html", cookieInfo)
}

func loginPage(res http.ResponseWriter, req *http.Request) {
    sessionCookie := retrieveCookie(res, req)
    if req.Method == "POST" && req.FormValue("password") == "csci130" {
        userInfo := decodeUserInformation(strings.Split(sessionCookie.Value, "|")[1])
        userInfo.Name = req.FormValue("username")
        userInfo.LoggedIn = true
        updateCookieValue(sessionCookie, userInfo)
        http.SetCookie(res, sessionCookie)
        http.Redirect(res, req, "/", http.StatusFound)
        return
    }
    tpl.ExecuteTemplate(res, "login.html", nil)
}
