package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todoapp "gin-todo-prc/src/models"
)

//	@Tags			lists
//	@Summary		Create List
//	@Description	Creating new todo list  for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			payload			body		todoapp.TodoList	true	"list info"
//	@Param			list_id			path		int					true	"list_id"
//	@Success		201				{string}	string				"list_id"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	var (
		userID, listID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	var input todoapp.TodoList
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if listID, err = h.services.TodoList.Create(userID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"list_id": listID,
	})
}

//	@Tags			lists
//	@Summary		Get all lists
//	@Description	Get all todo lists of current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Success		200				{onject}    todoapp.GetAllListsResponse
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	var (
		userID int
		lists  []todoapp.TodoList
		err    error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if lists, err = h.services.TodoList.GetAll(userID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todoapp.GetAllListsResponse{
		Data: lists,
	})
}

//	@Tags			lists
//	@Summary		Get list by Id
//	@Description	Get list by a specific id for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			list_id			path		int					true	"list_id"
//	@Success		200				{object}	todoapp.TodoItem
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 	   {object} errorResponse
//
//	@Router			/api/lists/{id} [get]
func (h *Handler) getListByID(c *gin.Context) {
	var (
		userID, id int
		list       todoapp.TodoList
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if list, err = h.services.TodoList.GetByID(userID, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

//	@Tags			lists
//	@Summary		Update a List
//	@Description	Updating a List by a its id for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			payload			body		todoapp.UpdateListInput	true	" update item info"
//	@Param			list_id			path		int					true	"list_id"
//	@Success		200				{string}	string				"status"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	var (
		userID, id int
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todoapp.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(userID, id, &input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

//	@Tags			lists
//	@Summary		Delete a list
//	@Description	Deleting a list by a its id for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			list_id			path		int					true	"list_id"
//	@Success		200				{string}	string				"status"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	var (
		userID, id int
		err        error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Delete(userID, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
