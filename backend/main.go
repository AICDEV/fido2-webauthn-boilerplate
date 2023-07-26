package main

import (
	"fmt"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("graceful terminate app. reason: '%s' ", sig.String())
		os.Exit(1)
	}()

	app.StartApplication()
}
