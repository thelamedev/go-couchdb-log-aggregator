package pkg

import "github.com/BurntSushi/toml"

type Config struct {
	App     AppConfig     `toml:"app"`
	CouchDB CouchDBConfig `toml:"couchdb"`
}

type AppConfig struct {
	Name        string `toml:"name"`
	Description string `toml:"desc"`
	Addr        string `toml:"addr"`
}

type CouchDBConfig struct {
	Url string `toml:"url"`
	Db  string `toml:"db"`
}

func NewConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
