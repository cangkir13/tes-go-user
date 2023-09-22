package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/cangkir13/tes-go-user/lib/core/model"
	"github.com/cangkir13/tes-go-user/lib/util"
	"github.com/gin-gonic/gin"
)

// Get list users
//	@Tags			users
//	@Summary		get list users
//	@Description	get listing users with limit offset
//	@ID				get-list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.Response
//	@Failure		404	{object}	util.Response
//	@Failure		401	{object}	util.Response
//	@Failure		500	{object}	util.Response
//	@Router			/users [get]
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

// Get one users
//	@Tags			users
//	@Summary		get one users
//	@Description	get oneing users with limit offset
//	@ID				get-one
//	@Param			id	path	int	true	"id user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.Response
//	@Failure		404	{object}	util.Response
//	@Failure		401	{object}	util.Response
//	@Failure		500	{object}	util.Response
//	@Router			/users/{id} [get]
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

// create user
//	@Tags			users
//	@Summary		create user
//	@Description	create user
//	@ID				create
//	@Accept			json
//	@Param			raw	body	CreateProps	true	"create new user return id"
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.Response
//	@Failure		404	{object}	util.Response
//	@Failure		401	{object}	util.Response
//	@Failure		500	{object}	util.Response
//	@Router			/users [post]
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

// update user
//	@Tags			users
//	@Summary		update user
//	@Description	update user
//	@ID				update
//	@Accept			json
//	@Param			id	path	int			true	"id user"
//	@Param			raw	body	UpdatePorps	true	"update user return id"
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.Response
//	@Failure		404	{object}	util.Response
//	@Failure		401	{object}	util.Response
//	@Failure		500	{object}	util.Response
//	@Router			/users/{id} [PUT]
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

// delete user
//	@Tags			users
//	@Summary		delete user
//	@Description	delete user
//	@ID				delete
//	@Accept			json
//	@Param			id	path	int	true	"id user"
//	@Produce		json
//	@Success		200	{object}	util.Response
//	@Failure		400	{object}	util.Response
//	@Failure		404	{object}	util.Response
//	@Failure		401	{object}	util.Response
//	@Failure		500	{object}	util.Response
//	@Router			/users/{id} [delete]
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
