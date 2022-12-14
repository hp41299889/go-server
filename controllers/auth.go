package controllers

import (
	"context"
	"go-server/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

// @Tags auth
// @Accept json
// @Produce json
// @Param JsonData body LoginBody true "des"
// @Success 200 {string} ok
// @Router /api/v0/login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		log.Println(ctx)
		var loginBody LoginBody
		defer cancel()

		if err := c.BindJSON(&loginBody); err != nil {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": loginBody.Username}})
	}
}
