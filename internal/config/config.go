package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Host  string `yaml:"Host"`
	MinIO struct {
		EndPoint        string `yaml:"EndPoint"`
		AccessKeyID     string `yaml:"AccessKeyID"`
		SecretAccessKey string `yaml:"SecretAccessKey"`
		UseSSL          bool   `yaml:"UseSSL"`
	} `yaml:"MinIO"`
	Database struct {
		Driver      string `yaml:"Driver"`
		Source      string `yaml:"Source"`
		MaxOpen     int    `yaml:"MaxOpen"`
		MaxIdle     int    `yaml:"MaxIdle"`
		MaxLifetime int    `yaml:"MaxLifetime"`
	} `yaml:"Database"`
}

func LoadCfg() Config {
	cfgPath := os.Getenv("OssHelperConfigPath")
	if cfgPath == "" {
		panic("请设置系统变量 OssHelperConfigPath，指定配置文件的位置！")
	}

	viper.SetConfigFile(cfgPath)

	if e := viper.ReadInConfig(); e != nil {
		panic(e.Error())
	}

	var cfg Config
	if e := viper.Unmarshal(&cfg); e != nil {
		panic(e)
	}
	return cfg
}
