package shared

import "github.com/spf13/viper"

type Config struct {
	PortApp string `mapstructure:"PORT_APP"`
	DBHost  string `mapstructure:"DB_HOST"`
	AppName string `mapstructure:"APP_NAME"`
	GoEnv   string `mapstructure:"GO_ENV"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return
}
