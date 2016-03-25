package main

import (
        "github.com/nu7hatch/gouuid"
        "log"
    )

func generateUUID() (string) {
    uuid, err := uuid.NewV4()
    if err != nil {
        log.Fatalln(err)
    }
    return uuid.String()
}
