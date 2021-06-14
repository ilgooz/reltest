package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ilgooz/reltest/app"
	"github.com/ilgooz/reltest/cmd/reltestd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
