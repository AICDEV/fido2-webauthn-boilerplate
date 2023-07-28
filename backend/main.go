package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/app"
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
