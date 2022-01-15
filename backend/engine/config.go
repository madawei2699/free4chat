package engine

import (
	"io/ioutil"

	"github.com/MixinNetwork/mixin/logger"
	"github.com/pelletier/go-toml"
)

type Configuration struct {
	Engine struct {
		Interface string `toml:"interface"`
		Address   string `toml:"address"`
		LogLevel  int    `toml:"log-level"`
		PortMin   uint16 `toml:"port-min"`
		PortMax   uint16 `toml:"port-max"`
	} `toml:"engine"`
	Turn struct {
		Host   string `toml:"host"`
		Secret string `toml:"secret"`
	} `toml:"turn"`
	RPC struct {
		Port int `toml:"port"`
	} `toml:"rpc"`
}

func Setup(path string) (*Configuration, error) {
	logger.Printf("Setup(%s)\n", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Configuration
	err = toml.Unmarshal(f, &conf)
	return &conf, err
}
