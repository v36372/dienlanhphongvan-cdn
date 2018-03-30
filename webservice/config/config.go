package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	StorageRoleSlave bool
	Web              Web
	Log              Log
	Resource         Resource
	Storage          Service
	Slave            Service
	Pot              Service
	Imgx             Service
	Compressor       Cmd
	Convertor        Cmd
	NewRelic         NewRelic
	WhiteList        []string
}

type Web struct {
	Host  string
	Port  string
	Debug bool
}

type Log struct {
	Dir        string
	LevelDebug bool
}

type Resource struct {
	RootDir string
}

type Service struct {
	Address string
}

type NewRelic struct {
	Enable  bool
	Name    string
	License string
}

type Cmd struct {
	Enable bool
	Exec   string
}

func Load() Config {
	// parse custom config
	ret := Config{}
	if err := viper.Unmarshal(&ret); err != nil {
		panic(err)
	}
	return ret
}
