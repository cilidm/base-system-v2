package core

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	once sync.Once
	Conf = &conf{}
)

type conf struct {
	DB     DBConf
	Server Server
	Redis  RedisConf
	App    SettingConf
	Zaplog ZapLogConf
}

type Server struct {
	HTTPPort     int
	ReadTimeout  int
	WriteTimeout int
}

type DBConf struct {
	DBType string
	DBUser string
	DBPwd  string
	DBHost string
	DBName string
	//DBPath string
}

type RedisConf struct {
	RedisAddr string
	RedisPWD  string
	RedisDB   int
}

type SettingConf struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  int
	WriteTimeout int
	PageSize     int
	JwtSecret    string
	UploadTmpDir string
	ImgSavePath  string
	ImgUrlPath   string
}

type ZapLogConf struct {
	Level         string `json:"level" yaml:"level"`
	Format        string ` json:"format" yaml:"format"`
	Prefix        string ` json:"prefix" yaml:"prefix"`
	Director      string ` json:"director"  yaml:"director"`
	LinkName      string ` json:"linkName" yaml:"link-name"`
	ShowLine      bool   ` json:"showLine" yaml:"showLine"`
	EncodeLevel   string ` json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `json:"logInConsole" yaml:"log-in-console"`
}

func InitConfig(tomlPath string) {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(tomlPath)
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal("read config failed: %v", err)
		}
		v.Unmarshal(&Conf)
	})
}
