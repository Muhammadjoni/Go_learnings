package handler

import (
	todoapp "gin-todo-prc/src/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Tags			auth
//	@Summary		sign up
//	@Description	register a new user
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		todoapp.User	true	"account info"
//	@Success		200		{integer}	integer			1
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var (
		input todoapp.User
		id    int
		err   error
	)

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if id, err = h.services.Authorization.CreateUser(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//	@Tags			auth
//	@Summary		sign in
//	@Description	logging in the user
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		todoapp.SignInInput	true	"username password"
//	@Success		200		{string}	string		token	"someToken"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var (
		input todoapp.SignInInput
		token string
		err   error
	)

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if token, err = h.services.Authorization.GenerateToken(input.Username, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
