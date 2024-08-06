/*
Copyright © 2024 NAME HERE yuchitan@kkcompany.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// linecountCmd represents the linecount command
var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Long: `Print line count of file
Example:
futil linecount -f /path/to/file`,

	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := cmd.Flags().GetString("filePath")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("linecount called for " + filePath)
	},
}

func init() {
	rootCmd.AddCommand(linecountCmd)

	linecountCmd.Flags().StringP("filePath", "f", "", "File path(required)")
	linecountCmd.MarkFlagRequired("filePath")
}
