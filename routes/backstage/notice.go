package backstage

import (
	"encoding/json"
	"gindemo01/public/sql"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	tsgutils "github.com/typa01/go-utils"
	"strconv"
	"time"
)

type data struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Order     int    `json:"order,omitempty"`
	BeginTime string `json:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

func POSTNotice(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(Claimadmins)
	var d data
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(201, gin.H{
			"err": "上传参数错误",
		})
		return
	}
	if d.Title == "" || d.Content == "" || d.BeginTime == "" || d.EndTime == "" {
		c.JSON(201, gin.H{
			"err": "上传参数错误",
		})
		return
	}
	s := noticetostring(d)
	var U sql_struct.Affairs
	U.AffairsId = tsgutils.GUID()
	U.Uid = Cla.UserId
	U.AffairsType = "notice"
	U.AffairsData = s
	U.State = "1"
	U.CreateBy = Cla.UserId
	U.UpdateBy = Cla.UserId
	U.CreateTime = time.Now().String()[0:19]
	U.UpdateTime = time.Now().String()[0:19]
	err := sql.AffairsAdd(U)
	if err != nil {
		c.JSON(201, gin.H{
			"msg": "上传失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "上传成功",
	})
}
func GETNotice(c *gin.Context) {
	D, _ := sql.AffairsFind("affairs_type = ?", "notice")
	if len(D) == 0 {
		c.JSON(200, gin.H{
			"err": "没有公告",
		})
		return
	}
	var list []struct {
		Data data   `json:"data"`
		Uuid string `json:"uuid"`
	}
	for _, v := range D {
		var d data
		err := json.Unmarshal([]byte(v.AffairsData), &d)
		if err != nil {
			c.JSON(201, gin.H{
				"err": "解析错误",
			})
			return
		}
		var l struct {
			Data data   `json:"data"`
			Uuid string `json:"uuid"`
		}
		l.Uuid = v.AffairsId
		l.Data = d
		list = append(list, l)
	}
	c.JSON(200, gin.H{
		"msg":  "获取成功",
		"list": list,
	})
}
func PUTNotice(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(Claimadmins)
	var d data
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(201, gin.H{
			"err": "上传参数错误",
		})
		return
	}
	if d.Title == "" || d.Content == "" || d.BeginTime == "" || d.EndTime == "" {
		c.JSON(201, gin.H{
			"err": "上传参数错误",
		})
		return
	}
	s := noticetostring(d)
	var U sql_struct.Affairs
	U.AffairsId = c.Param("uuid")
	U.AffairsData = s
	U.UpdateBy = Cla.UserId
	U.UpdateTime = time.Now().String()[0:19]
	err := sql.ReviseAffairs(U)
	if err != nil {
		c.JSON(201, gin.H{
			"msg": "上传失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "上传成功",
	})
}
func DELETENotice(c *gin.Context) {
	var uuid []string
	if err := c.ShouldBind(&uuid); err != nil {
		c.JSON(201, gin.H{
			"err": "上传参数错误",
		})
		return
	}
	var es []string
	for _, v := range uuid {
		if v == "" {
			c.JSON(201, gin.H{
				"err": "上传参数错误",
			})
			return
		}
		err := sql.DelAffaris(v)
		if err != nil {
			es = append(es, err.Error())
		}
	}
	if len(es) != 0 {
		c.JSON(201, gin.H{
			"msg": "删除失败",
			"err": es,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}
func noticetostring(d data) string {
	s := `{`
	s += `"title":"` + d.Title + `",`
	s += `"content":"` + d.Content + `",`
	s += `"order":` + strconv.Itoa(d.Order) + `,`
	s += `"begin_time":"` + d.BeginTime + `",`
	s += `"end_time":"` + d.EndTime + `"`
	s += `}`
	return s
}
