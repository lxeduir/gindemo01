package main

import (
	"fmt"
	"gindemo01/common"
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/public/sql"
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
	err := sql.InitDB()
	//fmt.Println(common.MysqlInfo.Id)
	if err != nil {
		fmt.Println("初始化失败 - 1")
		return //初始化数据库失败
	}
	err = redis.InitClirnt()
	if err != nil {
		fmt.Println("初始化失败 - 2")
		panic(err)
		return //初始化redis失败

	}
	//public.Online()
	fmt.Println("初始化成功")
}
