package conf

import (
	"github.com/freezeChen/studio-library/conf"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/util"
	"github.com/micro/go-config"
)

type Config struct {
	Name           string
	Version        string
	Env            string
	TCPPort        string
	WebSocketPort  string
	TCPKeepalive   bool
	TCPReadBuffer  int
	TCPWriteBuffer int
	Mysql *mysql.Config
	Redis *redis.Config
	Kafka string
}

func Init() (*Config, error) {
	var (
		Conf = &Config{}
	)
	cfg := config.NewConfig()

	source := conf.LoadFileSource(util.GetCurrentDirectory() + "/conf.yaml")
	cfg.Load(source)
	if err := cfg.Scan(Conf); err != nil {
		return nil, err
	}

	return Conf, nil
}
