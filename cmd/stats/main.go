package main

import (
	"github.com/yuseinishiyama/stats/pkg/providers/email"
	"github.com/yuseinishiyama/stats/pkg/writer"
)

func main() {
	_ = providers.Run()
	writer.Run()
}