package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func JsonLog() {
	log.Printf("Hello World")

	// info以上のログのみ出力されるように設定
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// それぞれのレベルで出力
	log.Error().Msg("errorで出力")
	log.Info().Msg("infoで出力")
	log.Debug().Msgf("debugで出力: %v", "ただし出力されない")
}
