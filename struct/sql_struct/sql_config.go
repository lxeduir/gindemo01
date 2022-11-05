package sql_struct

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
	//DeletedAt   string
}

type Admininfo struct {
	Uid      string
	Email    string
	Username string
	Passwd   string
	Token    string
	State    int
	//DeletedAt string
}
type UserImg struct {
	Uid        string
	Name       string
	File       string
	Updatatime string
}
