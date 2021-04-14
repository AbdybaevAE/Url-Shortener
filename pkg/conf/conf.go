package conf

import "github.com/spf13/viper"

type Config struct {
	DbUrl                 string `mapstructure:"DB_URL"`
	Addr                  string `mapstructure:"ADDRESS"`
	Algo                  string `mapstructure:"ALGORITHM"`
	RedisAddr             string `mapstructure:"REDIS_ADDRESS"`
	AccessTokenSecret     string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret    string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiredAt  string `mapstructure:"ACCESS_TOKEN_EXPIRED_AT"`
	RefreshTokenExpiredAt string `mapstructure:"REFRESH_TOKEN_EXPIRED_AT"`
}

func Load(path string) (*Config, error) {
	var config *Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
