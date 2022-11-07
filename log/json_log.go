package main

import (
	buildinerrors "errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

func JsonLog() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	err := outer()

	// 出力形式の設定
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	// info以上のログのみ出力されるように設定
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	buildinerror := buildinerrors.New("errors")
	// それぞれのレベルで出力
	log.Error().Err(buildinerror).Msg("errorで出力")
	log.Error().Stack().Err(err).Msg("")
	log.Info().Msg("infoで出力")
	log.Debug().Msgf("debugで出力: %v", "debugモードでのみ出力される")
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}
