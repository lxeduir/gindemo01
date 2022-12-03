package upload

import (
	"fmt"
	"gindemo01/public"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func Img(c *gin.Context) {

	uid, err1 := c.GetQuery("uid")
	name, err2 := c.GetQuery("name")
	file, err := c.FormFile("file")
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".img" && fileExt != ".jpeg" && fileExt != ".jpg" && fileExt != ".7z" {
		c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "上传文件格式错误"})
		return
	}
	U := public.UserinfoFind("uid = ?"+
		"", uid)
	if len(U) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if !err1 || !err2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err1,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 上传文件到指定的目录
	fileName := uid + "_" + name + "_" + strconv.FormatInt(time.Now().Unix(), 10) + path.Ext(file.Filename)
	err3 := c.SaveUploadedFile(file, "./Assets/"+fileName)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err3.Error(),
		})
		return
	}
	var u sql_struct.UserImg
	u.Uid = uid
	u.File = "./Assets/" + fileName
	u.Updatatime = strconv.FormatInt(time.Now().Unix(), 10)
	u.Name = name
	msg := public.UserImgAdd(u)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' upload success.", file.Filename),
		"msg":     msg,
	})
}
