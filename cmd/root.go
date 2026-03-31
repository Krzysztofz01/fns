package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

	c := config.GetConfiguration()

	// NOTE: Preload singleton printer
	_ = printer.GetPrinter()

	if c.UseSearchAsDefaultCommand && fallbackToDefaultCommand(args) {
		args = append([]string{searchCmd.Use}, args[1:]...)
	} else {
		args = args[1:]
	}

	rootCmd.SetArgs(args)
	rootCmd.SetContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failure: %s\n", err)
		os.Exit(1)
	}

	cancel()
	os.Exit(0)
}

func fallbackToDefaultCommand(args []string) bool {
	cmd, _, err := rootCmd.Find(args[1:])
	if err != nil {
		return false
	}

	if cmd.Use != rootCmd.Use {
		return false
	}

	if errors.Is(cmd.Flags().Parse(args[1:]), pflag.ErrHelp) {
		return false
	}

	return true
}

var rootCmd = &cobra.Command{
	Use:   "fns",
	Short: "Fuzzy Note Search",
	Long:  "Fuzzy Note Search... for those who have lost their patience for browsing through notes",
}
