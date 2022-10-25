package public

const SqlUserId = "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

//const sqlUserId = "users:root@(192.168.31.160:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

// sqlUserId 数据库账号密码
type Userinfo struct {
	Uid         string
	Email       string
	Username    string
	Passwd      string
	Token       string
	Permissions string
	Userstatus  int
	Signtime    string
}

type Admininfo struct {
	Uid      string
	Email    string
	Username string
	Passwd   string
	Token    string
	State    int
}
type Usertoken struct {
	Uid             string
	Token           string
	Refreshtoken    string
	Updatetime      int64
	Expirationtime  int64
	Rtepirationtime int64
}
type Useremailtoken struct {
	Uid            string
	Email          string
	Token          string
	Updatetime     int64
	Expirationtime int64
	State          int
}
