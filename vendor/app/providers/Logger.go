package providers

import (
	"fmt"
	"os"
	"time"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Config *Config `inject:"config"`
}

func NewLogger() *Logger {
	fmt.Println("NewLogger")
	return &Logger{}
}

func (this *Logger)Register() (error) {
	key, _ := this.Config.GetKey("log", "dir")
	dir := key.Value()
	// 判断日志目录是否存在
	stat, err := os.Stat(dir)
	if err != nil {
		// 新建目录
		os.Mkdir(dir, 0755)
	} else {
		if stat.Mode() != 0755 {
			os.Chmod(dir, 0755)
		}
	}

	year := time.Now().Format("2006-01")
	path := fmt.Sprintf("%s/%s-%s.log", dir, "go.server", year)
	file, _ := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Infoln("LoggerProvider Register");

	return nil
}
