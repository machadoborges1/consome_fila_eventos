package configs

import "github.com/spf13/viper"

type conf struct {
	User        string `mapstructure:"ORACLE_USER"`
	Password    string `mapstructure:"ORACLE_PASSWORD"`
	Host        string `mapstructure:"ORACLE_HOST"`
	Port        int    `mapstructure:"ORACLE_PORT"`
	ServiceName string `mapstructure:"ORACLE_SERVICE_NAME"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
