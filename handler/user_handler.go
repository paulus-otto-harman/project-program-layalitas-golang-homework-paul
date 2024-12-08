package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"homework/domain"
	"homework/service"
	"net/http"
)

type UserHandler struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserHandler(service service.UserService, logger *zap.Logger) UserHandler {
	return UserHandler{service: service, logger: logger}
}

func (ctrl *UserHandler) Registration(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		responseError(c, "invalid request body", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Register(&user); err != nil {
		responseError(c, "server error", err.Error(), http.StatusInternalServerError)
		return
	}

	responseCreated(c, user, "user registered")
}
