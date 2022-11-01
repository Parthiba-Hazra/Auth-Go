package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// This struct hold all values that needed to initialize the routes
type Config struct {
	E *gin.Engine
}

// This struct hold service properties for handler can run
type Handler struct {
}

// This function intialize the routes
func CreateHandler(c *Config) {
	newH := &Handler{}

	newRoute := c.E.Group(os.Getenv("APP_API_URL"))

	newRoute.GET("/myAccount", newH.MyAccount)
	newRoute.POST("/signup", newH.SignUp)
	newRoute.POST("/signin", newH.SignIn)
	newRoute.POST("/signout", newH.SignOut)
	newRoute.POST("/tokens", newH.Tokens)
	newRoute.POST("/image", newH.UploadImg)
	newRoute.DELETE("/image", newH.DeleteImg)
	newRoute.PUT("/myDetails", newH.MyDetails)

}

// Show user's details
func (h *Handler) MyAccount(c *gin.Context) {
	c.JSON(http.StatusOK, "It's your details")
}

func (h *Handler) SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, "It's signup handler")
}

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, "It's signin handler")
}

func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, "It's signout handler")
}

func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, "It's token handler")
}

func (h *Handler) UploadImg(c *gin.Context) {
	c.JSON(http.StatusOK, "upload profile image")
}

func (h *Handler) DeleteImg(c *gin.Context) {
	c.JSON(http.StatusOK, "delete profile image")
}

func (h *Handler) MyDetails(c *gin.Context) {
	c.JSON(http.StatusOK, "update user's details")
}
