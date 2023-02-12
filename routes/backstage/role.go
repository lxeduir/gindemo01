package backstage

import (
	"encoding/json"
	"gindemo01/public/sql"
	"gindemo01/struct/sql_del_struct"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AdminPermission struct {
	RoleId         string              `json:"RoleId"`
	Name           string              `json:"Name"`
	PermissionJson listAdminPermission `json:"PermissionJson"`
	Description    string              `json:"Description"`
	CreateBy       string              `json:"CreateBy"`
	CreateTime     string              `json:"CreateTime"`
	UpdateBy       string              `json:"UpdateBy"`
	UpdateTime     string              `json:"UpdateTime"`
	Orders         int                 `json:"Orders"`
}

func GetRole(c *gin.Context) {
	role, errSql := sql.AdminRoleFind("role_id LIKE ?", "%")
	if errSql != nil {
		c.JSON(201, gin.H{
			"err": errSql.Error(),
		})
	}
	var R []AdminPermission
	for _, v := range role {
		b := []byte(v.PermissionJson)
		var m listAdminPermission
		err := json.Unmarshal(b, &m)
		if err != nil {
			c.JSON(201, gin.H{
				"msg": err.Error(),
			})
			return
		}
		var r AdminPermission
		r.RoleId = v.RoleId
		r.Name = v.Name
		r.Description = v.Description
		r.CreateBy = v.CreateBy
		r.CreateTime = v.CreateTime
		r.UpdateBy = v.UpdateBy
		r.UpdateTime = v.UpdateTime
		r.Orders = v.Orders
		r.PermissionJson = m
		R = append(R, r)
	}
	c.JSON(200, gin.H{
		"list": R,
	})
}
func PostRole(c *gin.Context) {
	var u AdminPermission
	var U sql_struct.AdminRole
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		s, errSql := sql.AdminRoleFind("role_id LIKE ?", "%")
		if errSql != nil {
			c.JSON(201, gin.H{
				"err": errSql.Error(),
			})
		}
		s1, errSql := sql.AdminRoleFind("name = ?", u.Name)
		if errSql != nil {
			c.JSON(201, gin.H{
				"err": errSql.Error(),
			})
		}
		if len(s1) > 0 {
			c.JSON(201, gin.H{
				"err": "角色名已存在",
			})
			return
		}
		Role, errSql := sql.AdminRoleFind("role_id LIKE ?", "%")
		if errSql != nil {
			c.JSON(201, gin.H{
				"err": errSql.Error(),
			})
		}
		mp := make(map[int]int)
		for _, v := range Role {
			i, _ := strconv.Atoi(v.RoleId)
			mp[i] = v.Orders
		}
		for i := 1; i < len(Role)+10; i++ {
			if _, ok := mp[i]; !ok {
				U.RoleId = strconv.Itoa(i)
				break
			}
		}
		U.Name = u.Name
		U.Description = u.Description
		U.CreateBy = "等待实现"
		U.CreateTime = time.Now().String()[0:19]
		U.UpdateBy = "等待实现"
		U.UpdateTime = time.Now().String()[0:19]
		U.Orders = len(s) + 1
		js := `{"list":[`
		for i, v := range u.PermissionJson.List {
			js = js + `{"type":"` + v.Type + `","type_id":"` + v.Id + `","permission":"` + v.P + `"}`
			if i != len(u.PermissionJson.List)-1 {
				js = js + ","
			}
		}
		js = js + `]}`
		U.PermissionJson = js
		if err := sql.AdminRoleAdd(U); err != nil {
			c.JSON(201, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "添加成功",
			})
		}

	}

}
func PutRole(c *gin.Context) {
	var u AdminPermission
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {

		msg := sql.ReviseAdminrole(u.ToAdminrole())
		role, _ := sql.AdminRoleFind("role_id LIKE ?", "%")
		var R []AdminPermission
		for _, v := range role {
			b := []byte(v.PermissionJson)
			var m listAdminPermission
			err := json.Unmarshal(b, &m)
			if err != nil {
				c.JSON(201, gin.H{
					"msg": err.Error(),
				})
				return
			}
			var r AdminPermission
			r.RoleId = v.RoleId
			r.Name = v.Name
			r.Description = v.Description
			r.CreateBy = v.CreateBy
			r.CreateTime = v.CreateTime
			r.UpdateBy = v.UpdateBy
			r.UpdateTime = v.UpdateTime
			r.Orders = v.Orders
			r.PermissionJson = m
			R = append(R, r)
		}
		c.JSON(200, gin.H{
			"msg":  msg,
			"list": R,
		})

	}
}
func DelRole(c *gin.Context) {
	var u sql_struct.AdminRole
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		var U sql_del_struct.AdminRole
		U.RoleId = u.RoleId
		msg := sql.DelAdminRole(U)
		role, _ := sql.AdminRoleFind("role_id LIKE ?", "%")
		c.JSON(200, gin.H{
			"msg":  msg,
			"list": role,
		})
	}
}

func (a AdminPermission) ToAdminrole() sql_struct.AdminRole {
	var A sql_struct.AdminRole
	A.Name = a.Name
	A.Description = a.Description
	A.UpdateBy = a.UpdateBy
	A.Orders = a.Orders
	A.RoleId = a.RoleId
	js := `{"list":[`
	for i, v := range a.PermissionJson.List {
		js = js + `{"type":"` + v.Type + `","type_id":"` + v.Id + `","permission":"` + v.P + `"}`
		if i != len(a.PermissionJson.List)-1 {
			js = js + ","
		}
	}
	js = js + `]}`
	A.PermissionJson = js
	return A
}
