package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupConfig(environment string) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config" + environment)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msg("Failed load configuration : " + err.Error())
		return
	}

	log.Info().Msg("Success load configuration...")
}
