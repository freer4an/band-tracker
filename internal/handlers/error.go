package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type myError struct {
	Status  int
	Message string
}

func ErrorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	var data myError
	data.Message = http.StatusText(status)
	fmt.Println(data.Message)
	data.Status = status
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	if err := temp.Execute(w, data); err != nil {
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
}
