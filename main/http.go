package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rabbit struct {
	Name string
	Body []byte
}

func (r *Rabbit) save() error {
	filename := r.Name + ".txt"
	return ioutil.WriteFile(filename, r.Body, 0600)
}

func load(name string) (*Rabbit, error) {
	filename := name + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Rabbit{Name: name, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Name, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	rabbit := &Rabbit{Name: "WhiteRabbit", Body: []byte("This is a normal rabbit. No Jokes.")}
	rabbit.save()
	fmt.Fprintf(w, "<h1>Saved!</h1>")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}