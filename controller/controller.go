package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test_members2_rest/db"
)

type Controller struct {}

func(Controller) GetAll(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"users": db.GetAll(),
	})
}

func (Controller) GetOne(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"user": db.GetOne(id),
	})
}

func (Controller) Insert(c *gin.Context)  {
	name := c.PostForm("name")
	role := c.PostForm("role")
	db.Insert(name, role)
	c.JSON(http.StatusCreated, gin.H{
		"status":"OK",
	})
}

func (Controller) Update(c *gin.Context)  {
	n:= c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	name := c.PostForm("name")
	role := c.PostForm("role")
	db.Update(id, name, role)
	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
	})
}

func (Controller) Delete(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}


	db.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
	})
}