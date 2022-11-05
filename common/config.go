package common

type configModel struct {
	Server *serverModel `yaml:"server"`
	Mysql  *mysqlModel  `yaml:"mysql"`
	jwt    *jwtModel    `yaml:"jwt"`
}

type serverModel struct {
	Port string `yaml:"port"` // server port
}
type mysqlModel struct {
	Id string `yaml:"id"`
}
type jwtModel struct {
	key string `yaml:"key"`
}
