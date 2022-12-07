package public

import (
	"encoding/json"
	"gindemo01/public/sql"
	"strconv"
)

type rAdminPermission struct {
	RoleId         int `json:"role_id" gorm:"column:role_id"`
	Description    string
	PermissionJson listAdminPermission
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}
type listAdminPermission struct {
	List []permission `json:"list"`
}
type permission struct {
	Type string `json:"type"`
	Id   string `json:"type_id"`
	P    string `json:"permission"`
}

func Makemps(p string) map[string]map[int]int {
	b := []byte(p)
	var m listAdminPermission
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	var map1 map[string]map[int]int
	map1 = make(map[string]map[int]int)
	for _, v := range m.List {
		Id, _ := strconv.Atoi(v.Id)
		P, _ := strconv.Atoi(v.P)
		if map1[v.Type] == nil {
			map1[v.Type] = make(map[int]int)
		}
		map1[v.Type][Id] = P
	}
	return map1
}
func Authmap(RoleId string) map[string]map[int]int {
	R := Getpermission("role_id = ?", RoleId)
	var map1 map[string]map[int]int
	map1 = make(map[string]map[int]int)
	if len(R) == 0 {
		return nil
	} else {

		for _, v := range R[0].PermissionJson.List {
			Id, _ := strconv.Atoi(v.Id)
			P, _ := strconv.Atoi(v.P)
			if map1[v.Type] == nil {
				map1[v.Type] = make(map[int]int)
			}
			map1[v.Type][Id] = P
		}
		return map1
	}
}
func Getpermission(query string, args string) []rAdminPermission {
	P := sql.AdminRoleFind(query, args)
	var R []rAdminPermission
	for _, v := range P {
		b := []byte(v.PermissionJson)
		var m listAdminPermission
		err := json.Unmarshal(b, &m)
		if err != nil {
			return nil
		}
		var r rAdminPermission
		r.CreateBy = v.CreateBy
		r.CreateTime = v.CreateTime
		r.UpdateBy = v.UpdateBy
		r.UpdateTime = v.UpdateTime
		r.PermissionJson = m
		r.Description = v.Description
		R = append(R, r)
	}
	return R
}
