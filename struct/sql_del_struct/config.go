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
	DeletedAt string `json:"DeletedAt,omitempty"`
}
type UserRedis struct {
	Uid            string
	Data           string
	Type           string
	CreateTime     string `gorm:"autoCreateTime"`
	ExpirationTime string
	DeletedAt      string `json:"DeletedAt,omitempty"`
}
type UserLog struct {
	Uid       string
	Src       string
	data      string
	time      string `gorm:"autoCreateTime"`
	DeletedAt string `json:"DeletedAt,omitempty"`
}
type AdminPath struct {
	RoutId     int
	Name       string
	Title      string
	Icon       string
	Path       string
	Component  string
	Super      string
	Permission int
	Ctime      string `gorm:"autoCreateTime"`
	Utime      string `gorm:"autoUpdateTime"`
	DeletedAt  string `json:"DeletedAt,omitempty"`
}
type AdminRole struct {
	RoleId      int
	Name        string
	Description string
	Ctime       string `gorm:"autoCreateTime"`
	Utime       string `gorm:"autoUpdateTime"`
	Permission  int
	Orders      int
	DeletedAt   string `json:"DeletedAt,omitempty"`
}
type AdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson []Permission `gorm:"type:json"`
	DeletedAt      string       `json:"DeletedAt,omitempty"`
}

type Permission struct {
	Type string `json:"type"`
	Id   string `json:"type_id"`
	P    string `json:"permission"`
}
