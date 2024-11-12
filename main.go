// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/book-online-shop/database"
// 	"github.com/book-online-shop/handler"
// 	"github.com/book-online-shop/repository"
// 	"github.com/book-online-shop/service"
// )

// func main() {
// 	db, err := database.InitDB()
// 	if err != nil {
// 		log.Fatal("Gagal connect database!!!!", err)
// 	} else {
// 		log.Println("Success connect to database !!!")
// 	}
// 	defer db.Close()

// 	orderRepo := repository.NewOrderRepository(db)
// 	orderService := service.NewOrderService(&orderRepo)
// 	orderHandler := handler.NewOrderHandler(orderService)

// 	paymentRepo := repository.NewPaymentRepository(db)
// 	paymentService := service.NewPaymentService(&paymentRepo)
// 	paymentHandler := handler.NewPaymentHandler(paymentService)

// 	http.HandleFunc("/payments", paymentHandler.CreatePaymentHandler)
// 	http.HandleFunc("/orders", orderHandler.CreateOrderHandler)

// 	log.Println("Server started at :8080")
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"log"
	"net/http"

	"github.com/book-online-shop/database"
	"github.com/book-online-shop/handler"
	"github.com/book-online-shop/logger"
	"github.com/book-online-shop/repository"
	"github.com/book-online-shop/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Inisialisasi koneksi database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database!")
	}
	defer db.Close()

	// Inisialisasi repositori, service, dan handler
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(&orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Membuat router baru menggunakan Chi
	r := chi.NewRouter()
	logger.InitLogger()

	// Apply middleware untuk seluruh router
	r.Use(middleware.Logger)    // Middleware untuk logging
	r.Use(middleware.Recoverer) // Middleware untuk menangani panic

	// Define routes untuk request POST
	// r.Post("/payments", paymentHandler.CreatePaymentHandler)
	r.Post("/orders", orderHandler.CreateOrderHandler)
	// r.With(middleware.LoggingMiddleware).Post()

	// Mulai server dengan router Chi
	log.Println("[INFO] Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("[ERROR] Error starting server:", err)
	}
}
