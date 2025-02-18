package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"net"
	"net/http"
	"strings"
	"time"

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

func tcpCommand() *cli.Command {
	return &cli.Command{
		Name:  "tcp",
		Usage: "checks the health of a tcp service",
		Flags: []cli.Flag{&cli.UintFlag{
			Name:    "timeout",
			Value:   500,
			Aliases: []string{"t"},
		}},
		Action: func(ctx context.Context, command *cli.Command) error {
			url := command.Args().First()
			if strings.Trim(url, " ") == "" {
				return cli.Exit("uri is required", 1)
			}
			dial, err := net.DialTimeout("tcp", url, time.Millisecond*time.Duration(command.Uint("timeout")))
			if err != nil {
				return cli.Exit(err.Error(), 2)
			}
			if dial != nil {
				defer dial.Close()
			}
			return nil
		},
	}
}

func main() {
	err := (&cli.Command{
		Name:     "healthcheck",
		Usage:    "cli tool to check health of services",
		Commands: []*cli.Command{httpCommand(), tcpCommand()},
	}).Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
