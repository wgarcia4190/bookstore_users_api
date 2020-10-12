package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wgarcia4190/bookstore_oauth_go/oauth"
	"github.com/wgarcia4190/bookstore_users_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_users_api/internal/services"
	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

// Create creates a new user
// POST /users
func Create(c *gin.Context) {
	var user users.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.UserService.Create(&user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(oauth.IsPublic(c.Request)))
}

// Get returns an user which id match with the :user_id parameter
// GET /users/:user_id
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := getUserId(c)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.ID {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

// Update updates an user which id match with the :user_id parameter
// PUT /users/:user_id
func Update(c *gin.Context) {
	var user users.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	userId, userErr := getUserId(c)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.Update(&user, userId, isPartial)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
}

// Delete removes an user from the database.
// DELETE /users/:user_id
func Delete(c *gin.Context) {
	userId, userErr := getUserId(c)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := services.UserService.Delete(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search return a list of users from the database
// DELETE /internal/users/search?status=true
func Search(c *gin.Context) {
	status := c.Query("status")

	userSlice, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, userSlice.Marshall(oauth.IsPublic(c.Request)))
}

// Login returns an user which email and password match
// POST /users/login
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)

		return
	}

	loggedUser, err := services.UserService.Login(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, loggedUser.Marshall(oauth.IsPublic(c.Request)))
}

// getUserId returns the user_id parameter from the gin.Context
func getUserId(c *gin.Context) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("invalid user id")
	}

	return userId, nil
}
