package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/printer"
)

var (
	Version string
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute(args []string) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unexpected failure: %s\n", err)
			os.Exit(1)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChannel
		cancel()
	}()

	// NOTE: Preload singletons (configuration and printer)
	_ = config.GetConfiguration()
	_ = printer.GetPrinter()

	rootCmd.SetArgs(args)
	rootCmd.SetContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failure: %s\n", err)
		os.Exit(1)
	}

	cancel()
	os.Exit(0)
}

var rootCmd = &cobra.Command{
	Use:   "fns",
	Short: "Fuzzy Note Search",
	Long:  "Fuzzy Note Search... for those who have lost their patience for browsing through notes",
}
