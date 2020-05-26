package TestMongoDB

import (
	//log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

var (
	Conf = &Config{}
)

type Config struct {
	TCP              *TCP     `toml:"tcp"`
	MongoDB          *MongoDB `toml:"mongo-db"`
	LogMongoDB       *MongoDB `toml:"log-mongo-db"`
	DataDir          string   `toml:"data-dir"`
	Gm               bool     `toml:"gm"`
	AuthLoginAddress string   `toml:"auth-login-address"`
	WebHttpApiPort   string   `toml:"web-http-api-port"`
	PprofPort        string   `toml:"pprof-port"`
}

type TCP struct {
	Port    string `toml:"port"`
	BufSize int64  `toml:"buf-size"`
}

type MongoDB struct {
	Host      string `toml:"host"`
	Port      string `toml:"port"`
	PoolLimit int64  `toml:"pool-limit"`
	Source    string `toml:"source"` // 认证数据库
	UserName  string `toml:"user-name"`
	PassWd    string `toml:"pass-wd"`
	TimeOut   int64  `toml:"time-out"`
	DbName    string `toml:"db-name"`
}

func (c *Config) Fix() {
	if c.TCP.BufSize <= 0 {
		c.TCP.BufSize = 1024 * 1024 * 64
	}
}

// Load loads config options from a toml file.
func (c *Config) Load(confFile string) error {
	_, err := toml.DecodeFile(confFile, c)
	if err != nil {
		return err
	}
	c.Fix()
	return nil
}

func (c *Config) LoadFromData(data string) error {
	_, err := toml.Decode(data, c)
	if err != nil {
		return err
	}
	c.Fix()
	return nil
}
