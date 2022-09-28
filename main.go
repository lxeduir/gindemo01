package main

import (
	"gindemo01/backstage"
	"gindemo01/front"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*") //配置html模板
	//r.GET("/", front.IndexHtml)
	//r.POST("/login", front.LoginAdmin)
	backstage.Main(r)
	front.Main(r)
	universal.ErrorApi(r)
	r.Run(":8000")

}
