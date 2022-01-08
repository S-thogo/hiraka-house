package main

import (
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//page := Page{"Hello World.", 10}
	//tmpl, err := template.ParseFiles("layout/test.html") // ParseFilesを使う
	tmpl, err := template.ParseFiles("layout/index.html") // ParseFilesを使う

	if err != nil {
		panic(err)
	}

	//err = tmpl.Execute(w,page)
	err = tmpl.Execute(w,nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler) // hello

	dir, _ := os.Getwd()
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir(dir + "/layout/"))))

	http.ListenAndServe(":8080", nil)
}