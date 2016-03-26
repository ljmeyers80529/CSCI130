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
    http.ListenAndServe(":8080", nil)
}

// process root / top page.
func rootPage(res http.ResponseWriter, req *http.Request) {
    sessionCookie := retrieveCookie(res, req)
    if req.Method == "POST" {
        updateCookie(req, sessionCookie)
    }
    http.SetCookie(res, sessionCookie)
    //
    cookieInfo := cookieInformation {
        OriginalCookieValue: sessionCookie.Value,
        OriginalData: decodeUserInformation(strings.Split(sessionCookie.Value, "|")[1]),
    }
    fields := strings.Split(sessionCookie.Value,"|")
    ui := decodeUserInformation(fields[1])
    ui.Name = ui.Name + "1234"
    fields[1],_ = encodeUserInformation(ui)
    sessionCookie.Value = fields[0] + "|" + fields[1] + "|" + fields[2]
    
    var cookieState = "False"
    if isCookieInformationValid(sessionCookie) {
        cookieState = "True"
    }
    //
        cookieInfo.AlteredCookieValue = sessionCookie.Value
        cookieInfo.AlteredData = decodeUserInformation(strings.Split(sessionCookie.Value, "|")[1])
        cookieInfo.State = cookieState
    
    
    tpl.ExecuteTemplate(res, "root.html", cookieInfo)
}
