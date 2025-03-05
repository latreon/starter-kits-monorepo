package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starter-kit",
	Short: "A CLI tool to generate framework starter kits",
	Long:  `A CLI tool that helps you generate starter kits for various frameworks like React, Next.js, and Vue.js.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
} 