package main

import (
	"log"
	"net/http"

	"github.com/book-online-shop/database"
	"github.com/book-online-shop/handler"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal connect database!!!!", err)
	} else {
		log.Println("Success connect to database !!!")
	}
	defer db.Close()

	// r := chi.NewRouter()

	// r.Use(middleware.Logger) // Log setiap permintaan
	// r.Use(middleware.Recoverer)

	// r.Route("/auth", func(auth chi.Router) {
	// 	auth.Get("/loginpage", handler.LoginPage) // GET /auth/login
	// 	// auth.Post("/login", handler.LoginHandler) // POST /auth/login
	// })
	http.HandleFunc("/login", handler.LoginPage)
	http.HandleFunc("/auth/login", handler.LoginHandler)
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))

	//go to dashboard
	http.HandleFunc("/dashboard", handler.DashboardPage)

	http.HandleFunc("/add-book", handler.AddBookPage)
	http.HandleFunc("/book-list", handler.BookListPage)
	http.HandleFunc("/order-list", handler.OrderListPage)
	http.HandleFunc("/discount-book", handler.DiscoutPage)
	http.HandleFunc("/logout", handler.LogoutPage)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
