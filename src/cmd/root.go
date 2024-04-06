package cmd

import (
	"os"
	"time"

	"github.com/Lemuren/teal/cli"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "teal ADDRESS [PORT]",
	Short: "client for the TELNET protocol",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		timeout, err := cmd.Flags().GetDuration("timeout")
		if err != nil {
			os.Exit(1)
		}
		// Default to port 21 if no port was specified.
		if len(args) == 1 {
			cli.CliLoop(args[0]+":21", timeout)
			// Otherwise use specified port.
		} else {
			cli.CliLoop(args[0]+":"+args[1], timeout)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Duration("timeout", 3*time.Second, "")
}
