package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/Duarte64/go-web-meli/pkg/web"
	"github.com/gin-gonic/gin"
)

type User struct {
	service users.Service
}

type UserModelDto struct {
	Name     string  `json:"name" binding:"required"`
	Lastname string  `json:"lastname" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
}

type UserPatchDto struct {
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description list users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response{data=[]users.User}
// @Router /users [get]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := c.service.GetAll()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		}

		if len(u) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, u, ""))
	}
}

// GetUser godoc
// @Summary Get user
// @Tags Users
// @Description get user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response{data=users.User}
// @Router /users/:id [get]
func (c *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}

		u, err := c.service.GetById(uint(idInt))
		if err != nil {
			var notFoundErr *users.NotFoundError
			if errors.As(err, &notFoundErr) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, notFoundErr.Error()))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, u, ""))
	}
}

// StoreUser godoc
// @Summary Store user
// @Tags Users
// @Description store user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body UserModelDto true "User to store"
// @Success 200 {object} web.Response{data=users.User}
// @Router /users/:id [put]
func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userDto UserModelDto
		if err := ctx.ShouldBindJSON(&userDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		user, err := c.service.Store(userDto.Name, userDto.Lastname, userDto.Email, userDto.Age, userDto.Height, userDto.Active)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, user, ""))
	}
}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Description update user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body UserModelDto true "User to update"
// @Success 201 {object} web.Response{data=users.User}
// @Router /users [post]
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userDto UserModelDto
		if err := ctx.ShouldBindJSON(&userDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}

		user, err := c.service.Update(uint(id), userDto.Name, userDto.Lastname, userDto.Email, userDto.Age, userDto.Height, userDto.Active)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, user, ""))
	}
}

// PatchUser godoc
// @Summary Patch user
// @Tags Users
// @Description patch user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body UserPatchDto true "Fields to update"
// @Success 200 {object} web.Response{data=users.User}
// @Router /users/:id [patch]
func (c *User) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userPatchDto UserPatchDto
		if err := ctx.ShouldBindJSON(&userPatchDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}

		user, err := c.service.Patch(uint(id), userPatchDto.Lastname, userPatchDto.Age)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

// deleteUser godoc
// @Summary Delete user
// @Tags Users
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 204
// @Router /users/:id [delete]
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}
		if err := c.service.Delete(uint(id)); err != nil {
			var notFoundErr *users.NotFoundError
			if errors.As(err, &notFoundErr) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, notFoundErr.Error()))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
