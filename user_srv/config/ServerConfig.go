package config

type ServerConfig struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json="mysql"`
}
