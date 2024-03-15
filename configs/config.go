package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var cfg *Conf

type Conf struct {
	DbDriver      string           `mapstructure:"DB_DRIVER"`
	DbHost        string           `mapstructure:"DB_HOST"`
	DbPort        string           `mapstructure:"DB_PORT"`
	DbUser        string           `mapstructure:"DB_USER"`
	DbPassword    string           `mapstructure:"DB_PASSWORD"`
	DbName        string           `mapstructure:"DB_NAME"`
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string           `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int              `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth `mapstructure:""`
}

func LoadConfig(path string) (*Conf, error) {
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
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
