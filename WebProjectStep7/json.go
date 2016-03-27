package main

import (
        "encoding/json"
        "encoding/base64"
        "crypto/sha256"
        "crypto/hmac"
        "io"
        "fmt"
        "log"
        )

var key string = "NotSoSecretKey"

// marshal user information, convert to base 64 and get HMAC verfication code.
func encodeUserInformation(userInfo UserInformation) (string, string) {
    encoded, err := json.Marshal(userInfo)
    if err != nil {
        log.Fatalln(err)
    }
    b64 := base64.URLEncoding.EncodeToString(encoded)
    return b64, hash(b64)
}

// encode user information to be used for data validation.
func hash(info string) string {
    hmacHash := hmac.New(sha256.New, []byte(key))
    io.WriteString(hmacHash, info)
    return fmt.Sprintf("%x", hmacHash.Sum(nil))
}

// decode base 64, json information back into noremal format.
func decodeUserInformation(value string) UserInformation {
    var data UserInformation
    
    b64, err := base64.URLEncoding.DecodeString(value)
    if err != nil {
        log.Fatalln(err)
    }
    err = json.Unmarshal([]byte(b64), &data)
    if err != nil {
        log.Fatalln(err)
    }
    return data
}
