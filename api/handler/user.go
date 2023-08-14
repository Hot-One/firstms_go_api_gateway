package handler

import (
	"net/http"

	"github.com/Hot-One/firstms_go_api_gateway/genproto/user_service"
	"github.com/Hot-One/firstms_go_api_gateway/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Create user godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Procedure json
// @Param user body user_service.UserCreate true "CreateUserRequest"
func (h *Handler) CreateUser(ctx *gin.Context) {
	var user user_service.UserCreate

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"massage": err.Error(),
		})
		return
	}

	resp, err := h.services.UserService().Create(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "GRPC Error",
			"massage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// GetByID user godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *Handler) GetByIdUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	if !helper.IsValidUUID(userId) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Invalid uuid",
			"massage": "Invalid uuid",
		})
		return
	}

	resp, err := h.services.UserService().GetById(ctx, &user_service.UserPrimaryKey{Id: userId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "GRPC ERROR",
			"massage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
