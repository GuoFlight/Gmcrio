package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	GConf ConfigFileStruct
)

//解析配置文件
func ParseConfig(pathConfFile string) {
	if _, err := toml.DecodeFile(pathConfFile, &GConf); err != nil {
		log.Fatal(err)
	}
}

//配置文件结构体
type ConfigFileStruct struct {
	Log struct {
		Level         string `toml:"level"`
		Path          string `toml:"path"`
		RotationCount uint   `toml:"rotationCount"`
	} `toml:"log"`
	Http struct {
		Port int `toml:"port"`
	} `toml:"http"`
	Timer struct {
		Interval int `toml:"interval"`
	} `toml:"timer"`
}
