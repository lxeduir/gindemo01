package common

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

// LoadConfigInformation load config information for application

func loadConfigInformation(configPath string) (err error) {

	var (
		filePath string

		wr string
	)

	if configPath == "" {

		wr, _ = os.Getwd()

		wr = path.Join(wr, "conf")

	} else {

		wr = configPath

	}

	WorkSpace = wr

	filePath = path.Join(WorkSpace, "config.yml")

	configData, err := os.ReadFile(filePath)

	if err != nil {

		fmt.Printf(" config file read failed: %s", err)

		os.Exit(-1)

	}

	err = yaml.Unmarshal(configData, &ConfigInfo)

	if err != nil {

		fmt.Printf(" config parse failed: %s", err)

		os.Exit(-1)

	}

	// server information
	MysqlInfo = ConfigInfo.Mysql
	ServerInfo = ConfigInfo.Server
	Jwtinfo = ConfigInfo.Jwt
	Redisinfo = ConfigInfo.Redis
	return nil

}
func ReadConf() (err error) {
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err = loadConfigInformation(*configPath)
	//fmt.Printf("%+v\n", ConfigInfo)
	if err != nil {
		return err
	}
	return nil
}
