package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/book-online-shop/model"
	"github.com/book-online-shop/service"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Max 10 MB file upload
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// id := common.GenerateUUID()

	// id := r.FormValue("id")
	// Get form values
	name := r.FormValue("name")
	isActiveStr := r.FormValue("is_active")
	isActive, err := strconv.ParseBool(isActiveStr)
	if err != nil {
		http.Error(w, "Invalid value for is_active", http.StatusBadRequest)
		return
	}

	// Handle file upload
	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Photo is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file content into byte slice
	photoData := make([]byte, 0)
	_, err = file.Read(photoData)
	if err != nil {
		http.Error(w, "Failed to read photo file", http.StatusInternalServerError)
		return
	}

	// payload := model.Payment{}
	// payload.ID = uuid.New().String()
	// Create payment model
	payment := model.Payment{
		ID:       uuid.New().String(),
		Name:     name,
		Photo:    photoData,
		IsActive: isActive,
	}

	// Call service to create payment
	if err := h.paymentService.CreatePayment(payment); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create payment: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Payment created successfully"))
}
