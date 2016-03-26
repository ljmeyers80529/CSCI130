package main

import (
        "net/http"
        "strconv"
        "log"
        "encoding/json"
        "encoding/base64"
        "crypto/sha256"
        "crypto/hmac"
        "io"
        "fmt"
        )

type UserInformation struct {
    Name string
    Age int
    }
    
var key string = "NotSoSecretKey"

// marshal user information using JSON. Returns information in base 64 string and the HMAC verification encoding.
func jsonEncodeInformation(req *http.Request) (string, string) {
    a, _ := strconv.Atoi(req.FormValue("age"))
    user := UserInformation {
        Name: req.FormValue("username"),
        Age: a,
    }
    encoded, err := json.Marshal(user)
    if err != nil {
        log.Fatalln(err)
    }
    base64Encoded := base64.URLEncoding.EncodeToString(encoded)
    return base64Encoded, hash(base64Encoded)
}

func hash(info string) string {
    hash := hmac.New(sha256.New, []byte(key))
    io.WriteString(hash, info)
    return fmt.Sprintf("%x", hash.Sum(nil))
}

func jsonDecodeInformation(value string) UserInformation {
    var data UserInformation
    
    if value != "" {
        base64Decoded, err := base64.URLEncoding.DecodeString(value)
        if err != nil {
            log.Fatalln(err)
        }
        err = json.Unmarshal([]byte(base64Decoded), &data)
        if err != nil {
            log.Fatalln(err)
        }
        return data
    } else {
        return UserInformation {
                                Name: "",
                                Age: 0,
                                }
    }
}
