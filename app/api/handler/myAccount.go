package handler

import (
	"log"
	"net/http"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/Parthiba-Hazra/auth-go/models/error"
	"github.com/gin-gonic/gin"
)

// Show user's details
func (h *Handler) MyAccount(c *gin.Context) {

	user, present := c.Get("user")

	if !present {
		log.Printf("Unable to get user details :( \n reason: %v", c)
		err := error.NewServerErr()
		c.JSON(err.ErrorStatus(), gin.H{
			"error": err,
		})

		return
	}

	id := user.(*models.User).UID

	u, err := h.UserService.Get(c, id)

	if err != nil {
		log.Printf("Unable to find the user: %v\n err: %v", id, err)
		e := error.NewErrNotFound("user", id.String())

		c.JSON(e.ErrorStatus(), gin.H{
			"error": e,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	}
}
