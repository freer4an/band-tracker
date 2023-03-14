package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if !valid_url(id) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	// re := regexp.MustCompile("/artist/(\\d+)$")
	// matches := re.FindStringSubmatch(r.URL.Path)
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Println("Template parsing error")
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, all_data[id-1])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func valid_url(id int) bool {
	return (id > 0 && id <= len(all_data))
}
