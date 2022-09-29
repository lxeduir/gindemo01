package main

import (
	"gindemo01/backstage"
	"gindemo01/front"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	backstage.Main(r)
	front.Main(r)
	universal.ErrorApi(r)
	r.Run(":8000")
}
