/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package gcc

import (
	"github.com/spf13/cobra"
)

// gccCmd represents the gcc command
var GccCmd = &cobra.Command{
	Use:   "gcc",
	Short: "The Blockchain Bar CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	// rootCmd.AddCommand(GccCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gccCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gccCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
