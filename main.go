package main

import (
	"fmt"
	"text/template"
	"time"

	// "io"
	"log"
	"net/http"
)

type Film struct {
	Title string
	Director string
}

func main() {
	fmt.Println("hello world")

	h1 := func (w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello World\n")
		// io.WriteString(w, r.Method)
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func (w http.ResponseWriter, r *http.Request) {
		// log.Print("HTMX request received")
		// log.Print(r.Header.Get("HX-Request"))

		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)

		// fmt.Println(title)
		// fmt.Println(director)
	}



	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
