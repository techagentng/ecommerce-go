package handlers

import (
	_"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/gorm"

	_ "ecommerce-boiler/internals/core/domain"
	"ecommerce-boiler/internals/core/ports"
	"ecommerce-boiler/pkg/utils"
)

type userHandler struct {
	UserService ports.IUserService
	logger         *log.Logger
	handlerName    string
}

var (
	result  utils.Result
)

// NewUserHandler function creates a new instance for company handler
func NewUserHandler(cs ports.IUserService, l *log.Logger, n string) ports.IUserHandler {
	return userHandler{
		UserService: cs,
		logger:         l,
		handlerName:    n,
	}
}

func (ch userHandler) GetUserByID(c *gin.Context) {
		return
}

func (ch userHandler) GetAllUser(c *gin.Context) {
	var query utils.Pagination
	if err := c.ShouldBindQuery(&query); err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	users, err := ch.UserService.GetAllUser(&query)
	if err != nil {
		ch.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(users, "Successfully fetched all users"))
}

func (ch userHandler) CreateUser(c *gin.Context) {
	return
}

func (ch userHandler) DeleteUser(c *gin.Context) {
	return
}

func (ch userHandler) UpdateUser(c *gin.Context) {
	return
}
