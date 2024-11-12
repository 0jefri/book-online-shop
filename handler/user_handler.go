package handler

import (
	"encoding/json"
	"net/http"

	"github.com/book-online-shop/library"
	"github.com/book-online-shop/model"
	"github.com/book-online-shop/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewHandlerUser(us service.UserService) UserHandler {
	return UserHandler{UserService: us}
}

// var templates = template.Must(template.ParseGlob("view/*.html"))

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	err = h.UserService.LoginService(&user)
	if err != nil {
		library.BadResponse(w, "Account Not Found")
		return
	}

	library.SuccessResponse(w, "Login success", user)

	// templates.ExecuteTemplate(w, "/login", )
}
