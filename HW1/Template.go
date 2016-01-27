
// larry meyers
// CSCI 130
// 26-Jan-2016

package main

import (
        "fmt" 
        "os" 
        "text/template"
       )

type person struct {
    Name string
    Age uint8
    IQ uint8
    Genius bool
}

func main() {
    p1 := person {
            Name : "Todd McLeod",
            Age : 20,
            IQ: 155,
    }
    p1.Genius = p1.IQ >= 150
    tpl, err := template.ParseFiles("tpl_html_image.html")
    if err != nil {
        fmt.Println(err)
    }
    err = tpl.Execute(os.Stdout, p1)
    if err != nil {
        fmt.Println(err)
    }
}
