package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rabbit struct {
	Name string `json:"name,omitempty"`
	Body string `json:"description,omitempty"`
}

func load(name string) (*Rabbit, error) {
	filename := name + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Rabbit{Name: name, Body: string(body)}, nil
}

func getRabbit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/rabbit/"):]
	rabbit, _ := load(title)
	writeJSON(rabbit, w)
}

func writeJSON(model interface{}, w http.ResponseWriter) {
	JSON, err := json.Marshal(model)
	if err != nil {
		fmt.Fprintf(w, "Failed to marshal object to json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(JSON)
}

func main() {
	http.HandleFunc("/rabbit/", getRabbit)
	http.ListenAndServe(":8080", nil)
}
