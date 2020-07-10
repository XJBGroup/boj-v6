package user

import "github.com/gin-gonic/gin"

type Service interface {
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
