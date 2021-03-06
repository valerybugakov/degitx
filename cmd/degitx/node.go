// MIT License. Copyright (c) 2020 CQFN
// https://github.com/cqfn/degitx/blob/master/LICENSE

package main

import (
	"context"
	"log"

	"cqfn.org/degitx/discovery"
	"cqfn.org/degitx/locators"
)

// Start DeGitX node or return error
func Start(
	ctx context.Context,
	node *locators.Node,
	disc discovery.Service,
) error {
	log.Printf("Starting %s", node)

	if err := disc.Start(ctx); err != nil {
		return err
	}

	return nil
}
