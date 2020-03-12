package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arielril/hexgo/internal/container/model"

	uController "github.com/arielril/hexgo/internal/controller/user"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	Serve() error
}

type httpHandler struct {
	UserController uController.UserController
}

func NewHttpServer(ctx *HandlerContext) HttpHandler {
	return &httpHandler{
		UserController: ctx.UserController,
	}
}

func (h *httpHandler) Serve() error {
	engine := gin.New()

	engine.Use(gin.Recovery())

	api := engine.Group("/")

	api.GET("/simpleHealthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"alive": true,
		})
	})

	api.POST("/v1/user", func(c *gin.Context) {
		var user *model.User

		rawData, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to parse data",
				"err":     err,
			})
			return
		}

		_ = json.Unmarshal(rawData, user)
		if user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create entity from the request",
				"err":     err,
			})
			return
		}

		err = h.UserController.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to insert user",
				"err":     err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": user.ID,
			"created": true,
		})
	})

	return engine.Run(":3000")
}
