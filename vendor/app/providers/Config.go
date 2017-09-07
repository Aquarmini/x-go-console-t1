package providers

import (
	"github.com/go-ini/ini"
	"fmt"
)

type Config struct {
	Cfg *ini.File
}

func NewConfig() *Config {
	fmt.Println("NewConfig")
	cfg, _ := ini.InsensitiveLoad("config.ini")
	config := &Config{Cfg:cfg}
	return config
}

func (this *Config)Init() {

}

func (this *Config)GetSection(id string) (*ini.Section, error) {
	return this.Cfg.GetSection(id)
}

func (this *Config)GetKey(sec string, key string) (*ini.Key, error) {
	section, _ := this.Cfg.GetSection(sec)
	return section.GetKey(key)
}


