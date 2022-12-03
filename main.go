package main

import (
	"fmt"
	"gindemo01/common"
	"gindemo01/public"
	"gindemo01/routes/backstage"
	"gindemo01/routes/front"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(public.Cors())
	backstage.BackMain(r)
	front.Main(r)
	public.ErrorApi(r)
	r.Static("/assets", "./assets")
	fmt.Println("服务启动成功")
	r.Run(":" + common.ServerInfo.Port)

}
func init() {
	if common.ReadConf() != nil {
		return //读取配置文件失败
	}
	err := public.InitDB()
	if err != nil {
		return //初始化数据库失败
	}
	err = public.InitClirnt()
	if err != nil {
		return //初始化redis失败
	}
	fmt.Println("初始化成功")

}
