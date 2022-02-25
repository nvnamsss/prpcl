package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	Prefix = ""
)

var Config AppConfig

type AppConfig struct {
	Host    string `default:"0.0.0.0" envconfig:"HOST"`
	Port    int    `default:"8080" envconfig:"PORT"`
	RunMode string `default:"debug" envconfig:"RUN_MODE"`
	Env     string `default:"debug" envconfig:"ENV"`
	MySQL   MySQL
}

// AddressListener returns address listener of HTTP server.
func (c *AppConfig) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

func New() (*AppConfig, error) {
	if err := envconfig.Process(Prefix, &Config); err != nil {
		return nil, err
	}
	return &Config, nil
}

// MySQL represents configuration of MySQL database.
type MySQL struct {
	Username string `default:"root" envconfig:"MYSQL_USER"`
	Password string `default:"" envconfig:"MYSQL_PASS"`
	Host     string `default:"127.0.0.1" envconfig:"MYSQL_HOST"`
	Port     int    `default:"3306" envconfig:"MYSQL_PORT"`
	Database string `default:"prophet" envconfig:"MYSQL_DB"`
}

// ConnectionString returns connection string of MySQL database.
func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8"
	return fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database) + "&loc=Asia%2FHo_Chi_Minh"
}
