package config

import (
  "fmt"
  "github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Default defaultInfo
	Log logInfo
	Auth authInfo
}

type defaultInfo struct {
	Bind string
	Auth bool
}

type logInfo struct {
	File string
	Level string
	Console bool
}

type authInfo struct {
	Client_id string
	Client_key string
}

var Conf TomlConfig
func LoadConfig() {
	if _,err := toml.DecodeFile("f:\\work\\apiserver\\api.toml", &Conf); err != nil {
		fmt.Print(err)
		//PanicOnError(err)
	}
}