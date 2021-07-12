package main

import (
	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/command/bot"
	"github.com/yuseinishiyama/stats/command/gentoken"
	"github.com/yuseinishiyama/stats/command/worker"
)

func main() {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "personal statistics",
	}

	cmd.AddCommand(bot.Command())
	cmd.AddCommand(gentoken.Command())
	cmd.AddCommand(worker.Command())

	cmd.Execute()
}
