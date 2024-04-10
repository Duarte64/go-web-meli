package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Duarte64/go-web-meli/internal/users"
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

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}

		if len(u) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(200, u)
	}
}

func (c *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "ID inválido",
			})
			return
		}

		u, err := c.service.GetById(uint(idInt))
		if err != nil {
			var notFoundErr *users.NotFoundError
			if errors.As(err, &notFoundErr) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": notFoundErr.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, u)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userDto UserModelDto
		if err := ctx.ShouldBindJSON(&userDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Paylad enviado não corresponde a um usuário",
			})
			return
		}

		user, err := c.service.Store(userDto.Name, userDto.Lastname, userDto.Email, userDto.Age, userDto.Height, userDto.Active)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userDto UserModelDto
		if err := ctx.ShouldBindJSON(&userDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "ID inválido",
			})
			return
		}

		user, err := c.service.Update(uint(id), userDto.Name, userDto.Lastname, userDto.Email, userDto.Age, userDto.Height, userDto.Active)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	}
}

func (c *User) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userPatchDto UserPatchDto
		if err := ctx.ShouldBindJSON(&userPatchDto); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "ID inválido",
			})
			return
		}

		user, err := c.service.Patch(uint(id), userPatchDto.Lastname, userPatchDto.Age)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "ID inválido",
			})
			return
		}
		if err := c.service.Delete(uint(id)); err != nil {
			var notFoundErr *users.NotFoundError
			if errors.As(err, &notFoundErr) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": notFoundErr.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
