package handler

import (
	restapi "REST_API"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input restapi.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
