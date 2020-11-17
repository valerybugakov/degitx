// MIT License. Copyright (c) 2020 CQFN
// https://github.com/cqfn/degitx/blob/master/LICENSE

package main

import (
	"fmt"
	"log"
	"os"

	"cqfn.org/degitx/discovery"
	"cqfn.org/degitx/internal/config"
	"cqfn.org/degitx/logging"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:        "degitx",
		Usage:       "Manage degitx node",
		Description: "DeGitX node CLI",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "configuration file path",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run the node",
				Action:  cmdRun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "peer-host",
						Usage: "peer discovery host address (multiaddr)",
					},
					&cli.StringFlag{
						Name:  "peer-seed",
						Usage: "initial peer seed host address (multiaddr)",
					},
					&cli.StringFlag{
						Name:     "gitaly-host",
						Usage:    "Gitaly gRPC API host and port",
						Required: true,
					},
				},
			},
			{
				Name:   "id",
				Usage:  "print current node id",
				Action: printID,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("App failed: %s", err)
	}
}

const pConfigUser = "${HOME}/.config/degitx/config.yml"

const pConfigSys = "/etc/degitx/config.yml"

func cmdRun(ctx *cli.Context) error {
	cfg := new(config.DegitxConfig)
	if err := cfg.FromFiles(pConfigUser, pConfigSys, ctx.String("config")); err != nil {
		return err
	}
	node, err := cfg.Node()
	if err != nil {
		return err
	}
	logging.Init(node, cfg.LogConfig)
	peers := discovery.NewPeers(ctx.Context)
	var dsc discovery.Service
	peer := ctx.String("peer-host")
	seed := ctx.String("peer-seed")
	if peer != "" { //nolint:nestif,gocritic // consider refactoring later
		addr, err := ma.NewMultiaddr(peer)
		if err != nil {
			return err
		}
		dsc, err = discovery.NewGrpcServer(addr, node, peers)
		if err != nil {
			return err
		}
	} else if seed != "" {
		addr, err := ma.NewMultiaddr(seed)
		if err != nil {
			return err
		}
		dsc = discovery.NewGrpcClient(addr, node, peers)
	} else {
		dsc = new(discovery.StubService)
	}
	node.Addr, err = ma.NewMultiaddr(ctx.String("gitaly-host"))
	if err != nil {
		return err
	}

	return Start(ctx.Context, node, dsc)
}

func printID(ctx *cli.Context) error {
	cfg := new(config.DegitxConfig)
	if err := cfg.FromFiles(pConfigUser, pConfigSys, ctx.String("config")); err != nil {
		return err
	}
	l, err := cfg.Node()
	if err != nil {
		return err
	}
	fmt.Println(l)
	return nil
}