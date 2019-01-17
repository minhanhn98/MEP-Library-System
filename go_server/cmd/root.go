package cmd

import (
	"fmt"
	"mep-lib-system/internal/web"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goserve",
	Short: "goserve is a simple server for the MEP lending library system",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			logrus.Fatalf("could not get flag %q: %v\n", "port", err)
		}

		web.StartServer("localhost", port)

	},
}

func init() {
	rootCmd.Flags().IntP("port", "p", 8080, "the port to start the server on")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
