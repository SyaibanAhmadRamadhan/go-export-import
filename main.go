package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gpostgre"

	"github.com/SyaibanAhmadRamadhan/go-export-import/conf"
	"github.com/SyaibanAhmadRamadhan/go-export-import/repository"
	"github.com/SyaibanAhmadRamadhan/go-export-import/usecase"
)

func main() {
	err := genv.Initialize(genv.DefaultEnvLib, false)
	gcommon.PanicIfError(err)

	db := gpostgre.OpenPgxPool(conf.EnvPostgresConf().ConnString())
	defer db.Close()

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
		errs := bussines.LeaderBoard(context.Background(), os.Args[2:])
		if errs != nil {
			fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(errs.Error()))
			os.Exit(1)
		}
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Invalid command: %s\n", command)
		os.Exit(1)
	}
}
