package main

import (
	"github.com/yuseinishiyama/stats/command"
)

func main() {
	rootCmd := command.Command()
	rootCmd.Execute()
}
