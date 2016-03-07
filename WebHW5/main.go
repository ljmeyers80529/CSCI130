// web page homework #3
// larry meyere
// 06-Mar-2016

package main

import (
		"net/http"
		"io"
		"html/template"
		"log"
		"os"
		"path/filepath"
		"io/ioutil"
		)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	var txt string
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
		}
	if req.Method == "POST" {
		src, _, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join("./", "file.txt"))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
			}
		defer dst.Close()
		io.Copy(dst, src)
		bs, err := ioutil.ReadFile("file.txt")
		txt = string(bs)
		err = tpl.Execute(res, txt)
	} else {
			err = tpl.Execute(res, txt)
			if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
