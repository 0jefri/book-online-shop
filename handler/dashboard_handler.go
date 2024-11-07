// handler/dashboard.go
package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// DashboardPage menampilkan halaman dashboard setelah login berhasil
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "dashboard.html") // Path ke template dashboard
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
