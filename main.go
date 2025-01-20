package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"net/http"
	"strings"

	"os"
)

func httpCommand() *cli.Command {
	return &cli.Command{
		Name:  "http",
		Usage: "checks the health of an http service",
		Flags: []cli.Flag{&cli.UintFlag{
			Name:    "status",
			Value:   200,
			Aliases: []string{"c"},
		}},
		Action: func(ctx context.Context, command *cli.Command) error {
			url := command.Args().First()
			if strings.Trim(url, " ") == "" {
				return cli.Exit("url is required", 1)
			}
			get, err := http.Get(url)
			if err != nil {
				return cli.Exit(err.Error(), 2)
			}
			statusCode := command.Uint("status")
			if get.StatusCode != int(statusCode) {
				return cli.Exit(fmt.Sprintf("expected status code %d but got %d", statusCode, get.StatusCode), 1)
			}
			return nil
		},
	}
}

func main() {
	err := (&cli.Command{
		Name:     "healthcheck",
		Usage:    "cli tool to check health of services",
		Commands: []*cli.Command{httpCommand()},
	}).Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
