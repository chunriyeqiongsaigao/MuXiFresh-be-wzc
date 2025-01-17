package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string, prefix string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(prefix); err != nil {
		return err
	}

	return nil
}

func (c *Config) initConfig(prefix string) error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		// absPath, _ := filepath.Abs()
		viper.AddConfigPath("./config") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config.yaml")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量
	viper.SetEnvPrefix(prefix)  // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}
