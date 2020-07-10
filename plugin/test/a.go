package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Request struct {
	A int `form:"a" json:"a" binding:"required"`
	B int `form:"b" json:"b" binding:"required"`
}

type ErrorSerializer struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

const (
	CodeOK = iota
	CodeBindError
)

func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  CodeBindError,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var req = new(Request)
		if !BindRequest(c, req) {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":   CodeOK,
			"result": req.A + req.B,
		})
	})
	if err := r.Run(":12345"); err != nil {
		log.Fatalf("err serve %v", err)
	}
}
