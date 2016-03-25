package main

import (
        "net/http"
        "strings"
    )

var sessionCookie *http.Cookie
var sessionCookieName string = "session-fino"

//get an existing cookie or create "bake" a new cookie.
func eatCookie(res http.ResponseWriter, req *http.Request) (*http.Cookie) {
    myCookie, err := req.Cookie(sessionCookieName)
    if err != nil {
        myCookie = bakeNewCookie(res)
    }
    return myCookie
}

// create a new cookie, set the initial value with an UUID
func bakeNewCookie(res http.ResponseWriter) *http.Cookie {
    newCookie := &http.Cookie {
        Name: sessionCookieName,
        Value: generateUUID(),
        //Secure: true,
        HttpOnly: true,
        }
    http.SetCookie(res, newCookie)
    return newCookie
}

// update user information within the cookie. the user's uuid should always be the first element.
func rebakeCookie(myCookie *http.Cookie, req *http.Request) {
    var uuid string = myCookie.Value
    if strings.Contains(uuid, "|") {
        uuid = strings.Split(uuid, "|")[0]
    }
    jsonEncoded, hash := jsonEncodeInformation(req)
    myCookie.Value = uuid + "|" + string(jsonEncoded + "|" + hash)
}

// partition cookie into it three fields uuid, user inforation, verification value.
func partitionCookie(myCookie *http.Cookie) (string, string, string) {
    parts := strings.Split(myCookie.Value, "|")
    return parts[0], parts[1], parts[2]
}

// check if the cookie user information is valid.
func validCookie(cook *http.Cookie) string {
    fields := strings.Split(cook.Value, "|")
    bs := []byte(fields[1])
    bs[5] = bs[5]+1
    bs[6] = bs[6]-1
    bs[7] = bs[7]+2
    fields[1] = string(bs)
    cook.Value = fields[0]+"|"+fields[1]+"|"+fields[2]
    valueHash := hash(fields[1])
    if valueHash == fields[2] {
        return "Valid"
    } else {
        return "Invalid"
    }
}
