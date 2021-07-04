package main

import (
	"github.com/yuseinishiyama/stats/provider"
	"github.com/yuseinishiyama/stats/storage"
)

func main() {
	_ = provider.Run()
	storage.Run()
}