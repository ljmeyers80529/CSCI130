
// serfer web page
// larry meyere
// 01-Feb-2016

package main

import (
		"net/http";
		"html/template";
		"log"
		)

func surf(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}
func main() {
	http.Handle("/info/", http.StripPrefix("/info", http.FileServer(http.Dir("assests"))))
	http.HandleFunc("/", surf)
	http.ListenAndServe(":8080", nil)
}
 