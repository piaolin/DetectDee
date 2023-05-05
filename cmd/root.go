package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "DetectDee",
	Short: "Hunt down social media accounts by username, email or phone across social networks && Credential Stuffing them",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
