package providers

import (
	"github.com/go-ini/ini"
)

type Config struct {
	cfg *ini.File
}

func (this *Config)init() {
	if this.cfg == nil {
		this.cfg, _ = ini.InsensitiveLoad("config.ini")
	}
}
func (this *Config)GetSection(id string) (*ini.Section, error) {
	this.init()
	return this.cfg.GetSection(id)
}

func (this *Config)GetKey(sec string, key string) (*ini.Key, error) {
	this.init()
	section, _ := this.cfg.GetSection(sec)
	return section.GetKey(key)
}

