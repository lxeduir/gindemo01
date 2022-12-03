package sql_struct

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
	RoleId   int
	//DeletedAt string
}
type UserImg struct {
	Uid        string
	Name       string
	File       string
	Updatatime string `gorm:"autoUpdateTime"`
}
type UserRedis struct {
	Uid            string
	Data           string
	Type           string
	CreateTime     string `gorm:"autoCreateTime"`
	ExpirationTime string
}
type UserLog struct {
	Uid  string
	Src  string
	data string
	time string `gorm:"autoCreateTime"`
}
type AdminRout struct {
	RoutId     int
	Name       string
	Title      string
	Icon       string
	Component  string
	Path       string
	Super      string
	Permission int
	Ctime      string `gorm:"autoCreateTime"`
	Utime      string `gorm:"autoUpdateTime"`
}
type AdminRole struct {
	RoleId         string
	Name           string
	Description    string
	CreateBy       string
	CreateTime     string `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     string `gorm:"autoUpdateTime"`
	PermissionJson string
	Orders         int
}
type AdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson string
	CreateBy       string
	CreateTime     string `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     string `gorm:"autoUpdateTime"`
}

type Permission struct {
	Type string `json:"type"`
	Id   string `json:"type_id"`
	P    string `json:"permission"`
}
