package web

import (
	"os"

	"github.com/BurntSushi/toml"
)

//Load 加载配置文件
func Load(file string, obj interface{}) error {
	_, err := toml.DecodeFile(file, obj)
	return err
}

//Store 写入配置文件
func Store(file string, obj interface{}) error {
	fi, err := os.Create(file)
	defer fi.Close()

	if err == nil {
		end := toml.NewEncoder(fi)
		err = end.Encode(obj)
	}
	return err
}
