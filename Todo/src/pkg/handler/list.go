package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todoapp "gin-todo-prc/src/models"
)

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

type getAllListsResponse struct {
	Data []todoapp.TodoList `json:"data"`
}

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

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

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
