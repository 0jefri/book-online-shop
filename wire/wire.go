package wire

import (
	"github.com/book-online-shop/database"
	"github.com/book-online-shop/handler"
	"github.com/book-online-shop/repository"
	"github.com/book-online-shop/service"
	"github.com/google/wire"
	// "github.com/google/wire"
	// "github.com/book-online-shop/wire"
	// "github.com/google/wire"
)

var paymentHandlerSet = wire.NewSet(
	database.InitDB,
	repository.NewPaymentRepository,
	service.NewPaymentService,
	handler.NewPaymentHandler,
)

func InitializPaymentHandler() *handler.PaymentHandler {
	wire.Build(paymentHandlerSet)
	return nil
}
