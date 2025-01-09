package main

import (
	"buf-protoc/backend/component/config"
	"buf-protoc/backend/server"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func newApp() *server.Server {
	log.Debug().Str("app", "init").Msg("Initializing server...")
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
