package main

import (
	"os"

	"github.com/zhaotong0312/cloud-pubsub-beat/cmd"

	_ "github.com/zhaotong0312/cloud-pubsub-beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
