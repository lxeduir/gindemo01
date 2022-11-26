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
type AdminPath struct {
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
	RoleId      int
	Name        string
	Description string
	Ctime       string `gorm:"autoCreateTime"`
	Utime       string `gorm:"autoUpdateTime"`
	Permission  int
	Orders      int
}
type AdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson []Permission `gorm:"type:json"`
}

type Permission struct {
	Type string `json:"type"`
	Id   string `json:"type_id"`
	P    string `json:"permission"`
}
