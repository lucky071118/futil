/*
Copyright © 2024 NAME HERE yuchitan@kkcompany.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

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

		count, err := CountLine(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(count)
	},
}

func init() {
	rootCmd.AddCommand(linecountCmd)

	linecountCmd.Flags().StringP("filePath", "f", "", "File path(required)")
	linecountCmd.MarkFlagRequired("filePath")
}

func CountLine(filePath string) (int, error) {
	count := 0
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count += 1
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}
