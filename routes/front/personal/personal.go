package personal

import (
	"encoding/json"
	"gindemo01/public/sql"
	"gindemo01/routes/front/token"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	tsgutils "github.com/typa01/go-utils"
	"time"
)

type personal struct {
	Uid       string  `json:"uid"`
	About     string  `json:"about"`
	Applicant string  `json:"applicant"`
	Status    string  `json:"status"`
	EndTime   string  `json:"end_time"`
	Content   content `json:"content"`
}
type content struct {
	Isonechild string  `json:"isonechild"`
	Istwogirl  string  `json:"istwogirl"`
	Nperson    nperson `json:"nperson"`
	Cperson    cperson `json:"cperson"`
}
type nperson struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Doc     string `json:"doc"`
	Idcard  string `json:"idcard"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type cperson struct {
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Doc    string `json:"doc"`
	Idcard string `json:"idcard"`
	Phone  string `json:"phone"`
}
type data struct {
	About     string  `json:"about"`
	Applicant string  `json:"applicant"`
	Content   content `json:"content"`
}
type listdata struct {
	Datas data `json:"data"`
}

func PostPersonal(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	var p personal
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(201, gin.H{"err": "参数错误"})
	} else {
		var U sql_struct.Affairs
		u := sql.AffairsFind("uid = ?", Cla.UserId)
		U.AffairsId = tsgutils.GUID()
		U.Uid = Cla.UserId
		U.AffairsType = "personal"
		U.AffairsData = Datatostring(p)
		U.State = p.Status
		U.DisposeTime = ""
		U.UpdateBy = Cla.UserId
		U.UpdateTime = time.Now().String()[0:19]
		U.CreateBy = Cla.UserId
		U.CreateTime = time.Now().String()[0:19]
		for _, v := range u {
			if v.AffairsData == U.AffairsData {
				c.JSON(201, gin.H{"err": "请勿重复提交"})
				return
			}
		}
		sql.AffairsAdd(U)
		c.JSON(200, gin.H{"msg": "提交成功"})
	}
}
func GetPersonal(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	u := sql.AffairsFind("uid = ?", Cla.UserId)

	var ls []listdata

	for _, v := range u {
		if v.AffairsType == "personal" {
			b := []byte(v.AffairsData)
			var l listdata
			err := json.Unmarshal(b, &l)
			if err != nil {
				c.JSON(201, gin.H{"err": "数据解析错误"})
				return
			}
			ls = append(ls, l)
		}

	}
	c.JSON(200, gin.H{"datas": ls})
}
func Datatostring(p personal) string {
	var d data
	d.About = p.About
	d.Applicant = p.Applicant
	d.Content = p.Content
	s := `{"data":{`
	s += `"about":"` + d.About + `",`
	s += `"applicant":"` + d.Applicant + `",`
	s += `"content":{`
	s += `"isonechild":"` + d.Content.Isonechild + `",`
	s += `"istwogirl":"` + d.Content.Istwogirl + `",`
	s += `"nperson":{`
	s += `"name":"` + d.Content.Nperson.Name + `",`
	s += `"sex":"` + d.Content.Nperson.Sex + `",`
	s += `"doc":"` + d.Content.Nperson.Doc + `",`
	s += `"idcard":"` + d.Content.Nperson.Idcard + `",`
	s += `"phone":"` + d.Content.Nperson.Phone + `",`
	s += `"address":"` + d.Content.Nperson.Address + `"`
	s += `},`
	s += `"cperson":{`
	s += `"name":"` + d.Content.Cperson.Name + `",`
	s += `"sex":"` + d.Content.Cperson.Sex + `",`
	s += `"doc":"` + d.Content.Cperson.Doc + `",`
	s += `"idcard":"` + d.Content.Cperson.Idcard + `",`
	s += `"phone":"` + d.Content.Cperson.Phone + `"`
	s += `}`
	s += `}`
	s += `}}`
	return s
}
