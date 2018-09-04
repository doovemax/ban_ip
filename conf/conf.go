package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	// 根据操作系统选择配置文件
	switch runtime.GOOS {
	case "windows":
		ConfigFileName = "conf-windows.json"
	default:
		ConfigFileName = "conf.json"
	}

	ConfigPath = filepath.Dir(os.Args[0]) + PathSeparator + "conf.json"

	// 生成config配置文件
	logrus.Info("初始化配置文件:", ConfigPath)
	DefaultConfig, err = NewConfig(ConfigPath)

	if err != nil {
		logrus.Fatal(err)
	}
}

type Config struct {
	NginxFilePath string `json:"nginx_file_path"`
	LogFormat     string `json:"log_format"`
}

var (
	ConfigFileName string
	DefaultConfig  Config
	err            error
	PathSeparator  = string(os.PathSeparator)
	ConfigPath     string
)

// 生成新的配置文件
func NewConfig(filepath string) (Config, error) {
	fileByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(fileByte, &config)
	return config, err

}
