package config

type Cache struct {
	Address  string
	Password string
	Db       int
}

type Config struct {
	Cache Cache
}

func (cfg *Config) SetConfig(c Cache) {
	cfg.Cache.Address = c.Address
	cfg.Cache.Password = c.Password
	cfg.Cache.Db = c.Db
}

func (cfg *Config) GetConfig() Config {
	return *cfg
}
