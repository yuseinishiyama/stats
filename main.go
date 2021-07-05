// gen-token config
// put-token
// run

package main

import (
	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/command"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "stats",
		Run: func(cmd *cobra.Command, args []string) {
			command.Execute()
		},
	}

	rootCmd.Execute()
}
