package backstage

import (
	"gindemo01/public"
	"gindemo01/struct/sql_del_struct"
	"github.com/gin-gonic/gin"
)

type metas struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}
type super struct {
	Path      string  `json:"path"`
	Component string  `json:"component"`
	Name      string  `json:"name"`
	Meta      metas   `json:"meta"`
	Children  []super `json:"children,omitempty"`
}

func path(c *gin.Context) {
	cla, ok1 := c.Get("admininfo")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		paths := public.AdminRoutFind("rout_id LIKE ?", "%")
		var rpath []sql_del_struct.AdminRout
		for _, v := range paths {
			if Permissionvalidation(Cla.Mps, "rout", v.RoutId, 1) {
				rpath = append(rpath, v)
			}
		}
		rrpath := pathmake(rpath)
		c.JSON(200, gin.H{
			"uid":  Cla.UserId,
			"list": rrpath,
		})
	}
	return
}

func pathmake(path []sql_del_struct.AdminRout) interface{} {
	var rpath []super
	for _, v := range path {
		if v.Super == "/" {
			var m metas
			m.Title = v.Title
			m.Icon = v.Icon
			var s super
			s.Path = v.Path
			s.Component = v.Component
			s.Name = v.Name
			s.Meta = m
			rpath = append(rpath, s)
		}
	}
	for i, v := range rpath {
		var rpaths []super
		for _, vv := range path {
			if vv.Super == v.Path {
				var m metas
				m.Title = vv.Title
				m.Icon = vv.Icon
				var s super
				s.Path = vv.Path
				s.Component = vv.Component
				s.Name = vv.Name
				s.Meta = m
				rpaths = append(rpaths, s)
			}
		}
		rpath[i].Children = rpaths
	}
	return rpath
}
