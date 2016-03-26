package main

import (
        "net/http"
        "strings"
        )

var sessionCookieName string = "session-fino"

// get the requested cookie, create it it required.
func retrieveCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
    sessionCookie, err := req.Cookie(sessionCookieName)
    if err != nil {
        sessionCookie = createNewCookie()
    }
    return sessionCookie
}

// create a new cookie.
func createNewCookie() *http.Cookie {
    userInfo := defaultUserInformation()
    b64, hash := encodeUserInformation(userInfo)
    
    newSessionCookie := &http.Cookie {
        Name: sessionCookieName,
        Value: generateUUID() + "|" + b64 + "|" + hash,
        //Secure: true,
        HttpOnly: true,
    }
    return newSessionCookie
}

//update cookie information.
func updateCookie(req *http.Request, sessionCookie *http.Cookie) {
    userInfo := retrieveNewUserInformation(req)
    b64, hash := encodeUserInformation(userInfo)
    sessionCookie.Value = strings.Split(sessionCookie.Value,"|")[0] + "|" + b64 + "|" + hash
}

// validates cookie information has not been altered.
func isCookieInformationValid(sessionCookie *http.Cookie) bool {
    fields := strings.Split(sessionCookie.Value, "|")
    return hash(fields[1]) == fields[2]
}