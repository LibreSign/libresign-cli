package main

import (
	"fmt"
	"strings"

	"github.com/LibreSign/libresign-cli/pkg"
	"github.com/urfave/cli/v2"
)

var validate = &cli.Command{
	Name:  "validate",
	Usage: "validate PDF file",
	Action: func(c *cli.Context) error {
		input := strings.TrimSpace(c.Args().First())

		err := pkg.ValidatePDFFile(input)

		if err != nil {
			return cli.Exit(err, 1)
		}

		fmt.Println("valid document")

		return nil
	},
}
