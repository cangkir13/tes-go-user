package usercontroller

import (
	"strconv"

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

		c.JSON(400, result)
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

		c.JSON(400, result)
		return
	}

	result := util.Response{
		Status: false,
		Data:   user,
	}

	c.JSON(200, result)
}
