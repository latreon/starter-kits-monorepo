package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/latreon/starter-kits-monorepo/apps/cli/cmd"
	"github.com/latreon/starter-kits-monorepo/apps/cli/internal/utils"
)

func main() {
	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Handle interrupt in a separate goroutine
	go func() {
		<-sigChan
		utils.HandleInterrupt()
	}()

	cmd.Execute()
} 