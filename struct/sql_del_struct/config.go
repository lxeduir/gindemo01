package sql_del_struct

type Userinfo struct {
	Uid         string `json:"uid,omitempty"`
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	Passwd      string `json:"passwd,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Userstatus  int    `json:"userstatus,omitempty"`
	Signtime    string `json:"signtime,omitempty"`
	DeletedAt   string `json:"deletedAt,omitempty"`
}
type Admininfo struct {
	Uid       string
	Email     string
	Username  string
	Passwd    string
	Token     string
	State     int
	RoleId    int
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
type AdminRout struct {
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
	RoleId         string
	Name           string
	Description    string
	CreateBy       string
	CreateTime     string `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     string `gorm:"autoUpdateTime"`
	PermissionJson string
	Orders         int
	DeletedAt      string `json:"DeletedAt,omitempty"`
}
type AdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson string
	CreateBy       string
	CreateTime     string `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     string `gorm:"autoUpdateTime"`
	DeletedAt      string `json:"DeletedAt,omitempty"`
}
type Affairs struct {
	AffairsId   string
	Uid         string
	AffairsType string
	AffairsData string
	State       string
	DisposeTime string
	CreateBy    string
	CreateTime  string `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  string `gorm:"autoUpdateTime"`
	DeletedAt   string `json:"DeletedAt,omitempty"`
}
type UserIdentity struct {
	Uid       string
	Name      string
	Sex       int
	Ethnic    string
	Polstatus string
	Origin    string
	Value     string
	Docunum   string
	Birthdate string
	Phonenum  string
	DeletedAt string `json:"DeletedAt,omitempty"`
}
