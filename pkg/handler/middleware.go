package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPars := strings.Split(header, " ")
	if len(headerPars) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.services.Authorization.ParseToken(headerPars[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}
	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "invalid type user id")
		return 0, errors.New("invalid type user id")
	}
	return idInt, nil
}
