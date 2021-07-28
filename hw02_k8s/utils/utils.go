package utils

import (
	"fmt"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/mitchellh/mapstructure"
)

type DBConfig struct {
	Host string `mapstructure:"dbhost"`
	Port string `mapstructure:"dbport"`
	User string `mapstructure:"dbuser"`
	Pass string `mapstructure:"dbpass"`
	Name string `mapstructure:"dbname"`
}

func readDbConfig() (cfg DBConfig, err error) {
	input, err := beego.AppConfig.GetSection("db")
	if err != nil {
		return
	}
	err = mapstructure.Decode(input, &cfg)
	fmt.Println(err, cfg)

	return
}

func GetSQLDSN() (cfg string, err error) {
	c, err := readDbConfig()
	if err != nil {
		return
	}
	cfg = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", c.Host, c.Port, c.Name, c.User, c.Pass)

	return
}

func OptionalEmail(v *validation.Validation, obj interface{}, key string) {
	name, ok := obj.(string)
	if !ok {
		return
	}

	// valid true
	if len(name) == 0 {
		return
	}

	v.Email(obj, key)
}

func OptionalPhone(v *validation.Validation, obj interface{}, key string) {
	name, ok := obj.(string)
	if !ok {
		return
	}

	// valid true
	if len(name) == 0 {
		return
	}

	v.Phone(obj, key)
}
