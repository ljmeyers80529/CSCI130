package main

import (
    "net/http"
    "html/template"
    )

var tpl *template.Template

type CookieState struct {
    OriginalCookieValue string
    AlteredCookieValue string
    State string
    OriginalData UserInformation
    AlteredData UserInformation
    }

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
    var cookState CookieState
    
    myCookie := eatCookie(res, req)
    if req.Method == "POST" {
        rebakeCookie(myCookie, req)
    }
    //
    _, data, _ := partitionCookie(myCookie)
    origValue := myCookie.Value
    origUser := jsonDecodeInformation(data)
    state := validCookie(myCookie)                      // simulates altering the cookie before validating the cookie.
    alterValue := myCookie.Value;
    _, data, _ = partitionCookie(myCookie)
    alterUser := jsonDecodeInformation(data)
    //
    cookState = CookieState {
        OriginalCookieValue: origValue,
        State: state,
        AlteredCookieValue: alterValue,
        OriginalData: origUser,
        AlteredData: alterUser,
        }
    //
    tpl.Execute(res, cookState)
}
