package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "libresign",
		Usage:    "libresign portable cli tool",
		Commands: []*cli.Command{validate},
		ExitErrHandler: func(context *cli.Context, err error) {
			exiter, is := err.(cli.ExitCoder)
			if is {
				log.Error().Err(err).Msg("execution error")
				cli.OsExiter(exiter.ExitCode())
			}
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal().Err(err).Msg("execution failure")
	}
}
