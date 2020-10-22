package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	ConfigRuntime()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin()  {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/go/ping", ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}