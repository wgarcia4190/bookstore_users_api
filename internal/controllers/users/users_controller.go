package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wgarcia4190/bookstore_users_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_users_api/internal/services"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

// Create creates a new user
// POST /users
func Create(c *gin.Context) {
	var user users.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get returns an user which id match with the :user_id parameter
// GET /users/:user_id
func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
