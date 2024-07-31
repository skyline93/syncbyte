package api

import (
	"net/http"
	"phto/internal/auth"
	"phto/internal/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
//
//	@Summary		Login user
//	@Description	login user
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			user	body	UserLoginRequest	true	"user info"
//	@Router			/login [post]
func Login(router *gin.Engine) {
	handler := func(c *gin.Context) {
		var json UserLoginRequest

		if c.Bind(&json) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user := entity.FindUser(json.Username)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		if !user.InvalidPassword(json.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			return
		} else {
			token, exp := auth.GenerateToken(json.Username, user.Role)
			c.JSON(http.StatusOK, gin.H{"token": token, "exp": exp})
		}
	}

	router.POST("/login", handler)
}

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// Register godoc
//
//	@Summary		Register user
//	@Description	register user
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			user	body	RegisterUserRequest	true	"user info"
//	@Router			/register [post]
func Register(router *gin.Engine) {
	handler := func(c *gin.Context) {
		var json RegisterUserRequest

		if c.Bind(&json) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user := entity.FindUser(json.Username)
		if user != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
			return
		}

		role := auth.UserRoleUser
		if json.Role != "" && json.Role != auth.UserRoleAdmin {
			role = json.Role
		}

		if err := entity.CreateUser(json.Username, string(hashedPassword), role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User registered failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}

	router.POST("/register", handler)
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	delete user
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			username	path	string	true	"User Name"
//	@Router			/delete/:username [delete]
func DeleteUser(router *gin.Engine) {
	handler := func(c *gin.Context) {
		usernameToDelete := c.Param("username")
		currentUser, _ := c.Get("username")

		toDeletedUser := entity.FindUser(usernameToDelete)
		if toDeletedUser == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username not found"})
			return
		}

		user := entity.FindUser(currentUser.(string))
		if user == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username not found"})
			return
		}

		if user.Role != auth.UserRoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient privileges"})
			return
		}

		if err := entity.DeleteUser(usernameToDelete); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}

	router.DELETE("/delete/:username", auth.AuthMiddleware(), auth.AdminMiddleware(), handler)
}
