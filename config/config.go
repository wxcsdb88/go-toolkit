package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig load config
func LoadConfig(file string) (interface{}, error) {
	if "" == file {
		return nil, fmt.Errorf("blank file")
	}

	v := viper.New()
	v.SetConfigFile(file) // auto detect the file suffix

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c := new(interface{})
	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	return c, nil
}
