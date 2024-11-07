// handler/login.go
package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "login.html") // Path ke template login
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Fungsi untuk menangani login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validasi login sederhana
		if username == "jefri" && password == "123" {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			tmplPath := filepath.Join("templates", "login.html")
			tmpl, _ := template.ParseFiles(tmplPath)
			tmpl.Execute(w, map[string]string{"Error": "Invalid username or password"})
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
