package notice

import (
	"encoding/json"
	"fmt"
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type data struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Order     int    `json:"order,omitempty"`
	BeginTime string `json:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

func Get(c *gin.Context) {
	u, err := sql.AffairsFind("affairs_type = ?", "notice")
	if err != nil {
		c.JSON(201, gin.H{
			"err": err.Error(),
		})
		return
	}
	var list []data
	for _, v := range u {
		var d data
		err := json.Unmarshal([]byte(v.AffairsData), &d)
		st, _ := time.Parse("2006-01-02 15:04:05", d.BeginTime)
		et, _ := time.Parse("2006-01-02 15:04:05", d.EndTime)
		nt, _ := time.Parse("2006-01-02 15:04:05", time.Now().String()[0:19])
		fmt.Println("st:", st.Unix(), "et:", et.Unix(), time.Now().Unix())
		if err != nil {
			c.JSON(201, gin.H{
				"err": err.Error(),
			})
		}
		if nt.Unix() >= st.Unix() && nt.Unix() <= et.Unix() {
			list = append(list, d)
		}
	}
	c.JSON(200, gin.H{
		"list": list,
	})
}
