package controllers

import (
	"GolangProject/apis"
	"GolangProject/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AccountController interface {
	GetAccounts() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
	CreateAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
}

func NewAccountsController(service services.AccountService) AccountController {
	return &accountController{
		service: service,
	}
}

type accountController struct {
	service services.AccountService
}

func (a *accountController) GetAccounts() gin.HandlerFunc {
	return func(context *gin.Context) {
		acc, err := a.service.GetAccounts()
		if err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		context.JSON(http.StatusOK, acc)
	}
}

func (a *accountController) GetAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		acc, err := a.service.GetAccount(id)
		if err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		context.JSON(http.StatusOK, acc)
	}
}

func (a *accountController) CreateAccount() gin.HandlerFunc {
	return func(context *gin.Context) {

		var createRequest apis.AccountCreateRequest

		if err := context.ShouldBindJSON(&createRequest)
			err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		acc, err := a.service.CreateAccount(createRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, acc)
	}
}

func (a *accountController) UpdateAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		var updateRequest apis.AccountUpdateRequest

		id := context.Param("id")
		if err := context.BindJSON(&updateRequest)
			err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		acc, err := a.service.UpdateAccount(id, updateRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, acc)
	}
}

func (a *accountController) DeleteAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		if err := a.service.DeleteAccount(id); err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusNoContent, "Account deleted successfully.")
	}
}

func handleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.AbortWithError(statusCode, err)
}
