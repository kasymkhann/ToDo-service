package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	todo "to-doProjectGo"
)

func (h *Handler) createItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid list param")
		return
	}
	var input todo.ToDoItem
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.ToDoItem.Create(userId, listId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid list id param")
		return
	}

	items, err := h.services.ToDoItem.GetAll(userId, listId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
func (h *Handler) getItemsId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid list id param")
		return
	}
	item, err := h.services.ToDoItem.GetItemId(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
func (h *Handler) putItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	var input todo.UpdateItemListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ToDoItem.Update(userId, id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})

}
func (h *Handler) deleteItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid list id param")
		return
	}
	err = h.services.ToDoItem.Delete(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"OK"})
}
