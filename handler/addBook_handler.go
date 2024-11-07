package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func AddBookPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "add_book.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func BookListPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "book_list.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func OrderListPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "order_list.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DiscoutPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "discount_book.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("view", "logout.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
