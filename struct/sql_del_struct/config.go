package sql_del_struct

type Userinfo struct {
	Uid         string
	Email       string
	Username    string
	Passwd      string
	Token       string
	Permissions string
	Userstatus  int
	Signtime    string
	DeletedAt   string
}

type Admininfo struct {
	Uid       string
	Email     string
	Username  string
	Passwd    string
	Token     string
	State     int
	DeletedAt string
}
