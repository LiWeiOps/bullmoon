package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Model        string `mapstructure:"model"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Filename    string `mapstructure:"filename"`
	Level       string `mapstructure:"level"`
	Max_size    int    `mapstructure:"max_size"`
	Max_age     int    `mapstructure:"max_age"`
	Max_backups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	Dbname         string `mapstructure:"dbname"`
	Max_open_conns int    `mapstructure:"max_open_conns"`
	Max_idle_conns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Password  string `mapstructure:"password"`
	Db        int    `mapstructure:"db"`
	Pool_size int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config.yaml")         // 指定配置文件路径
	viper.AddConfigPath(".")                     // 还可以在工作目录中查找配置
	if err := viper.ReadInConfig(); err != nil { // 处理读取配置文件的错误
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("voper unmashai failed, err: %v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置发生变化")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("voper unmashai failed, err: %v\n", err)
		}
	})
	return
}
