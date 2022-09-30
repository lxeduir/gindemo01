package sql_operate

type Userinfo struct {
	Uid    string
	Name   string
	Passwd string
	Email  string
}
type Admininfo struct {
	Uid    string
	Name   string
	Passwd string
	Token  string
	Rtoken string
}
type Usertoken struct {
	Uid            string
	Token          string
	Updatetime     int64
	Expirationtime int64
}
type Useremailtoken struct {
	Uid            string
	Email          string
	Updatetime     int64
	Expirationtime int64
	State          int
	Token          string
}
