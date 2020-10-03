package server

import (
	"github.com/gin-gonic/gin"
	"test_members2_rest/controller"
)

func Init()  {
	r := router()
	r.Run("localhost:8080")
}

func router()  *gin.Engine{
	r := gin.Default()
	ctrl := controller.Controller{}

	// GET
	r.GET("/user", ctrl.GetAll)
	r.GET("/user/:id", ctrl.GetOne)

	// POST
	r.POST("/user", ctrl.Insert)

	// PUT
	r.PUT("/user/:id", ctrl.Update)

	// DELETE
	r.DELETE("/user/:id", ctrl.Delete)

	return r
}