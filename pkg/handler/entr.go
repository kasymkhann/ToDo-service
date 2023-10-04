package handler

import (
	"net/http"
	user "to-doProjectGo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorRespons(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Entering.CreateUser(input)
	if err != nil {
		NewErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignUp(c *gin.Context) {

}
