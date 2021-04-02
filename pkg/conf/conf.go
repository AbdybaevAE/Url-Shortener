package conf

import "github.com/spf13/viper"

type Config struct {
	DbUrl string `mapstructure:"DB_URL"`
	Addr  string `mapstructure:"ADDRESS"`
	Algo  string `mapstructure:"ALGORITHM"`
}

func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
