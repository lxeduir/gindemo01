package main

import (
	"gindemo01/common"
	"gindemo01/public"
	"gindemo01/routes/backstage"
	"gindemo01/routes/front"
	"github.com/gin-gonic/gin"
)

func main() {
	if common.ReadConf() != nil {
		return
	}
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(public.Cors())
	backstage.BackMain(r)
	front.Main(r)
	public.ErrorApi(r)
	r.Static("/assets", "./assets")
	r.Run(":" + common.ServerInfo.Port)

}
