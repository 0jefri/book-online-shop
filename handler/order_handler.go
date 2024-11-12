package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/book-online-shop/common"
	"github.com/book-online-shop/logger"
	"github.com/book-online-shop/model"
	"github.com/book-online-shop/service"
	"go.uber.org/zap"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Logger.Error("Invalid request method",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
		)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload model.Order
	payload.ID = common.GenerateUUID()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		logger.Logger.Error("Error decoding JSON",
			zap.String("handler", "CreateOrderHandler"),
			zap.String("function", "CreateOrderHandler"),
			zap.Error(err),
		)
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrder(&payload); err != nil {
		logger.Logger.Error("Error creating order",
			zap.String("handler", "CreateOrderHandler"),
			zap.String("function", "CreateOrderHandler"),
			zap.Error(err),
		)
		log.Println("Error creating order:", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Order created successfully",
		Data:       payload,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Logger.Error("Error encoding response",
			zap.String("handler", "CreateOrderHandler"),
			zap.String("function", "CreateOrderHandler"),
			zap.Error(err),
		)
		log.Println("Error encoding response:", err)
	}
}
