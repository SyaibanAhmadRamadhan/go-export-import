package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gpostgre"

	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/conf"
	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/repository"
	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/usecase"
)

func main() {
	err := genv.Initialize(genv.DefaultEnvLib, false)
	gcommon.PanicIfError(err)

	db := gpostgre.OpenPgxPool(conf.EnvPostgresConf().ConnString())
	postgres := gpostgre.NewPgxPostgres(db)
	competitionRepo := repository.NewCompetitionRepositoryImpl(postgres)
	matchRepo := repository.NewMatchRepositoryImpl(postgres)
	txRepo := gpostgre.NewTxPgx(db)

	bussines := usecase.NewCompetitionUsecaseImpl(competitionRepo, matchRepo, txRepo)

	command := os.Args[1]
	switch command {
	case "input":
		errs := bussines.Input(os.Args[2:], 1)
		if errs != nil {
			fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(errs.Error()))
			os.Exit(1)
		}
	case "leaderboard":
		leaderboard(os.Args[2:])
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Invalid command: %s\n", command)
		os.Exit(1)
	}
}

func leaderboard(args []string) {
	fmt.Println("leaderboard")
}
