package common

type configModel struct {
	Server *serverModel `yaml:"server"`
	Mysql  *mysqlModel  `yaml:"mysql"`
	Jwt    *jwtModel    `yaml:"jwt"`
}

type serverModel struct {
	Port string `yaml:"port"` // server port
	Host string `yaml:"host"` // server host
	SSL  bool   `yaml:"ssl"`  // server ssl
}
type mysqlModel struct {
	Id string `yaml:"id"`
}
type jwtModel struct {
	Key string `yaml:"key"`
}
