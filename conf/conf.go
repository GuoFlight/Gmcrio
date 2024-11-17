package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	GConf ConfigFile
)

// ConfigFile 配置文件结构体
type ConfigFile struct {
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
	Db struct {
		Server          string        `toml:"server"`
		MaxIdleConns    int           `toml:"maxIdleConns"`
		MaxOpenConns    int           `toml:"maxOpenConns"`
		ConnMaxLifetime time.Duration `toml:"connMaxLifetime"`
	} `toml:"db"`
	Jwt struct {
		Secret string `toml:"secret"`
		Expire int    `toml:"expire"`
	} `toml:"jwt"`
}

// ParseConfig 解析配置文件
func ParseConfig(pathConfFile string) {
	if _, err := toml.DecodeFile(pathConfFile, &GConf); err != nil {
		log.Fatal(err)
	}
	CheckAndInit()
}
func CheckAndInit() {
	// 设置工作目录
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	wd := filepath.Dir(path)
	fmt.Println("[INFO] work directory:", wd)
	err = os.Chdir(wd)
	if err != nil {
		log.Fatal(err)
	}
}
