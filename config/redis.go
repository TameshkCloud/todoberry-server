package config

type RedisConfig struct {
	Cluster      bool     `mapstructure:"cluster"`
	DSN          []string   `mapstructure:"dsn"`
	Timeouts     struct {
		InternalPoolTimeout int64 `mapstructure:"pool_timeout_in_seconds"`
		IdleTimeout         int64 `mapstructure:"idle_timeout_in_seconds"`
		ReadTimeout         int64 `mapstructure:"read_timeout_in_seconds"`
		WriteTimeout        int64 `mapstructure:"write_timeout_in_seconds"`
	} `mapstructure:"timeouts"`
	Pool struct {
		InitialCapacity int `mapstructure:"initial_capacity"`
		MaxCapacity     int `mapstructure:"max_capacity"`
	} `mapstructure:"connection_pool"`
}

var Redis RedisConfig

func InitRedisConfig() {
	LoadConfig("redis", &Redis)
}
