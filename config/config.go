package config

import (
	"github.com/spf13/viper"
)

var AppConfig = struct {
	Nacos NacosConfig
	Addr  string
}{
	Addr: ":8080",
}

type NacosConfig struct {
	Endpoint []NacosEndpoint
}
type NacosEndpoint struct {
	Addr string
}

func Load() error {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&AppConfig)
	return err
}
