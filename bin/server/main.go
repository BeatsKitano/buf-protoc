package main

import (
	"buf-protoc/backend/server"
	"buf-protoc/component/config"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (

	// greetingBanner is the greeting banner.
	// http://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=Bytebase
	greetingBanner = `
___________________________________________________________________________________________
██╗  ██╗██╗   ██╗██╗ ██████╗ ███████╗
██║  ██║██║   ██║██║██╔════╝ ██╔════╝
███████║██║   ██║██║██║  ███╗█████╗  
██╔══██║██║   ██║██║██║   ██║██╔══╝  
██║  ██║╚██████╔╝██║╚██████╔╝███████╗
╚═╝  ╚═╝ ╚═════╝ ╚═╝ ╚═════╝ ╚══════╝
___________________________________________________________________________________________
`
)

func newApp() *server.Server {
	return server.NewServer(":36789", &config.Profile{})
}

func main() {

	app := newApp()

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	// Trigger graceful shutdown on SIGINT or SIGTERM.
	// The default signal sent by the `kill` command is SIGTERM,
	// which is taken as the graceful shutdown signal for many systems, eg., Kubernetes, Gunicorn.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		slog.Info(fmt.Sprintf("%s received.", sig.String()))
		if app != nil {
			_ = app.Shutdown(ctx)
		}
		cancel()
	}()

	fmt.Println(string([]byte{27, 91, 51, 49, 109}), greetingBanner, string([]byte{27, 91, 48, 109}))
	fmt.Println(fmt.Sprintf("has started on port %d 🚀", 36789))

	// Execute program.
	app.Run()

	// Wait for CTRL-C.
	<-ctx.Done()
}
