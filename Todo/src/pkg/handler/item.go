package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todoapp "gin-todo-prc/src/models"
)

//	@Tags			items
//	@Summary		Create Item
//	@Description	Creating new item inside the list for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			payload			body		todoapp.TodoItem	true	"item info"
//	@Param			id				path		int					true	"id"
//	@Success		201				{string}	string				"id"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todoapp.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []todoapp.TodoItem
}

//	@Tags			items
//	@Summary		Get all items
//	@Description	Get all items from the list of current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			payload			body		todoapp.TodoItem	true	"item info"
//	@Param			id				path		int					true	"id"
//	@Success		201				{string}	string				"id"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 {object} errorResponse
//
//	@Router			/api/lists/:id/items [post]
func (h *Handler) getAllItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

//	@Tags			items
//	@Summary		Get item by Id
//	@Description	Get items Id for current user
//	@Accept			json
//	@Produce		json
//	@param			Authorization	header		string				true	"Bearer Auth, pls add bearer before"
//	@Param			payload			body		todoapp.TodoItem	true	"item info"
//	@Param			id				path		int					true	"id"
//	@Success		201				{string}	string				"id"
//
// Failure default {object} errorResponse
// Failure 400,404 {object} errorResponse
// Failure 500 	   {object} errorResponse
//
//	@Router			/api/items/:id [get]
func (h *Handler) getItemByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetByID(userID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	var (
		userID, itemID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if itemID, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	var input todoapp.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err = h.services.TodoItem.Update(userID, itemID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	var (
		userID, itemID int
		err            error
	)

	if userID, err = getUserID(c); err != nil {
		return
	}

	if itemID, err = strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err = h.services.TodoItem.Delete(userID, itemID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
