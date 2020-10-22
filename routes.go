package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ping(c *gin.Context) {
	time.Sleep(50 * time.Millisecond)
	c.JSON(http.StatusOK, gin.H{})
}
