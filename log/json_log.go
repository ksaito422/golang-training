package main

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func JsonLog() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	// info以上のログのみ出力されるように設定
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// それぞれのレベルで出力
	log.Error().Msg("errorで出力")
	log.Info().Msg("infoで出力")
	log.Debug().Msgf("debugで出力: %v", "debugモードでのみ出力される")
}
