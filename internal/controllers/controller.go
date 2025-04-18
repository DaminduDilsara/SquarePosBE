package controllers

import (
	"github.com/Square-POC/SquarePosBE/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerV1 struct {
	servicesCollection *services.ServiceCollection
}

func NewControllerV1(servicesCollection *services.ServiceCollection) *ControllerV1 {
	return &ControllerV1{
		servicesCollection: servicesCollection,
	}
}

// Route: Redirect to Google for authentication
func (con *ControllerV1) GoogleLogin(c *gin.Context) {
	url := con.servicesCollection.AuthSvc.OAuthLogin()
	c.Redirect(http.StatusFound, url)
}

func (con *ControllerV1) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	userInfo, err := con.servicesCollection.AuthSvc.OAuthCallBack(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OAuth error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Google Login Successful",
		"user":    userInfo,
	})
}
