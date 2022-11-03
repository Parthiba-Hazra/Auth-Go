package handler

import (
	"log"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/Parthiba-Hazra/auth-go/models/error"
	"github.com/gin-gonic/gin"
)

// Show user's details
func (h *Handler) MyAccount(c *gin.Context) {

	user, present := c.Get("user")

	if !present {
		log.Printf("Unable to get user details :( reason: %v", c)
		err := error.NewServerErr()
		c.JSON(err.ErrorStatus(), gin.H{
			"error": err,
		})

		return
	}

	id := user.(*models.User).UID
}
