package main

import (
	"fmt"
	_ "fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)
var (
	templateFile = template.Must(template.ParseFiles("templates/index.html"))
	port = 9090
)
func main() {
	http.HandleFunc("/", upload)
	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleUpload(w,r)
		return
	}
	templateFile.ExecuteTemplate(w, "index.html", nil)
}
func handleUpload(w http.ResponseWriter, r *http.Request) { 
	r.ParseMultipartForm(10 << 20) // 10 MB
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	defer file.Close()
	log.Printf("file name: %s", fileHeader.Filename)
	log.Printf("file size: %d", fileHeader.Size)
	filename := path.Base(fileHeader.Filename)
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, "INternal server error", http.StatusInternalServerError)
	}

	if _, err := io.Copy(dest, file); err != nil {
		http.Error(w, "INternal server error", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}