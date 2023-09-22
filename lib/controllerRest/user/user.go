package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	"github.com/cangkir13/tes-go-user/lib/util"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) GetListUser(c *gin.Context) {
	users, err := ctrl.userusecase.GetListUser(c)

	if err != nil {
		result := util.Response{
			Status: false,
			Error:  err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := util.Response{
		Status: true,
		Error:  "",
		Data:   users,
	}
	c.JSON(200, result)
}

func (ctrl *UserController) GetOneID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.userusecase.GetOneByID(c, id)
	if err != nil {
		result := util.Response{
			Status: false,
			Error:  err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := util.Response{
		Status: true,
		Data:   user,
	}

	c.JSON(200, result)
}

type CreateProps struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *UserController) Create(c *gin.Context) {
	var request CreateProps
	if err := c.ShouldBindJSON(&request); err != nil {
		result := util.Response{Status: false, Error: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	hashing, _ := util.HashPassword(request.Password)
	_, err := ctrl.userusecase.CreateOne(c, model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashing,
	})

	if err != nil {
		result := util.Response{
			Status: false,
			Error:  err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(200, util.Response{Status: true, Data: "data has been added"})
}

type UpdatePorps struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var request UpdatePorps
	if err := c.ShouldBindJSON(&request); err != nil {
		result := util.Response{Status: false, Error: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	err := ctrl.userusecase.UpdateOne(c, id, model.User{
		Name:     request.Name,
		Password: request.Password,
	})

	if err != nil {
		result := util.Response{
			Status: false,
			Error:  err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(200, util.Response{Status: true, Data: "data has been updated"})
}

func (ctrl *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.userusecase.DeleteOne(c, id)
	if err != nil {
		result := util.Response{
			Status: false,
			Error:  err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(200, util.Response{Status: true, Data: "data has been deleted"})
}
