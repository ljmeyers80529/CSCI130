package main

import (
        "strconv"
        "net/http"
        )

type UserInformation struct {
    Name string
    Age int
    LoggedIn bool
}

type cookieInformation struct {
    OriginalCookieValue string
    OriginalData UserInformation
    State string
}

// set user information to default values.
func defaultUserInformation() UserInformation {
    userInfo := UserInformation {
        Name: "",
        Age: 0,
        LoggedIn: false,
    }
    return userInfo
}

// retrive user information for user entered information on the web form.
func retrieveNewUserInformation(req *http.Request) UserInformation {
    age, _ := strconv.Atoi(req.FormValue("age"))
    userInfo := UserInformation {
        Name: req.FormValue("username"),
        Age: age,
        LoggedIn: true,
    }
    return userInfo
}