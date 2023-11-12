package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorRespons(c, http.StatusOK, "empty auth header")
		return
	}

	headerParts := strings.Split(header, "")
	if len(headerParts) != 2 {
		NewErrorRespons(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		NewErrorRespons(c, http.StatusUnauthorized, err.Error())

	}

	c.Set(userCtx, userId)
}
