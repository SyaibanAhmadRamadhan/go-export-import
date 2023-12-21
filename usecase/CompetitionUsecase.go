package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gdir"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/model"
	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/repository"
)

// data struct represents a combination of Competition and Match models.
type data struct {
	Compeititon model.Competition
	Match       model.Match
}

// CompetitionUsecaseImpl is the implementation of the CompetitionUsecase interface.
// It manages the business logic related to competitions and matches.
type CompetitionUsecaseImpl struct {
	compeRepo repository.CompetitionRepository
	matchRepo repository.MatchRepository
	tx        gdb.Tx
}

// NewCompetitionUsecaseImpl creates a new instance of CompetitionUsecaseImpl.
// It takes instances of CompetitionRepository, MatchRepository, and a transaction object as parameters.
func NewCompetitionUsecaseImpl(
	compeRepo repository.CompetitionRepository,
	matchRepo repository.MatchRepository,
	tx gdb.Tx,
) *CompetitionUsecaseImpl {
	return &CompetitionUsecaseImpl{
		compeRepo: compeRepo,
		matchRepo: matchRepo,
		tx:        tx,
	}
}

// Input method processes CSV file data, performs validation, and inserts records into the database.
// It uses a transaction to ensure data consistency.
// The 'limit' parameter controls the batch processing size for max paramter postgres, ensuring data is processed in manageable chunks.
func (u *CompetitionUsecaseImpl) Input(args []string, limit int) error {
	if len(args) == 0 {
		return errors.New("invalid arguments input")
	}

	dir, err := gdir.FindDirPathOfFileFromGoMod(args[0])
	if err != nil {
		return fmt.Errorf("cannot find file %s: %s", args[0], err)
	}

	gomodDir, err := gdir.LocateGoModDirectory()
	if err != nil {
		return fmt.Errorf("cannot find file go.mod, %s", err)
	}

	file, err := os.Open(gomodDir + "/" + dir + "/" + args[0])
	if err != nil {
		return fmt.Errorf("failed open file %s/%s: %s", dir, args[0], err)
	}

	reader := csv.NewReader(file)
	defer func() {
		if errClose := file.Close(); errClose != nil {
			log.Info().Msgf("failed close file reader, %s", errClose)
		}
	}()

	header := make([]string, 0)
	counter := 0
	ctx := context.Background()
	var datas []data

	err = u.tx.DoTransaction(ctx, &gdb.TxOption{
		Type: gdb.TxTypePgx,
		Option: pgx.TxOptions{
			IsoLevel:       pgx.Serializable,
			AccessMode:     pgx.ReadWrite,
			DeferrableMode: pgx.NotDeferrable,
		},
	}, func(c context.Context) (commit bool, err error) {
		for {
			rows, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					err = nil
				} else {
					return commit, err
				}
				break
			}

			if len(header) == 0 {
				header = rows
				continue
			}

			csvDataRow, err := u.fetchDataCsv(rows)
			if err != nil {
				if !errors.Is(err, ErrMustRollback) {
					commit = true
				}
				return commit, err
			}

			datas = append(datas, data{
				Compeititon: model.Competition{
					ID:   csvDataRow.CompetitionID,
					Date: csvDataRow.Date,
				},
				Match: model.Match{
					CompeititonID: csvDataRow.CompetitionID,
					Team1Name:     csvDataRow.Team1,
					Team1Score:    csvDataRow.Team1Score,
					Team2Name:     csvDataRow.Team2,
					Team2Score:    csvDataRow.Team2Score,
				},
			})
			counter++

			if counter%limit == 0 {
				newDatas := append([]data(nil), datas...)
				datas = nil
				err = u.doTheJob(c, newDatas)
				if err != nil {
					return commit, err
				}
			}
		}

		if len(datas) > 0 {
			err = u.doTheJob(c, datas)
			if err != nil {
				return commit, err
			}
		}

		return commit, err
	})

	return err
}

// doTheJob method processes the collected Data and inserts records into the database.
func (u *CompetitionUsecaseImpl) doTheJob(ctx context.Context, datas []data) (err error) {
	var competitions []model.Competition
	var matches []model.Match
	for _, v := range datas {
		competitions = append(competitions, v.Compeititon)
		matches = append(matches, v.Match)
	}

	err = u.compeRepo.InsertMany(ctx, competitions)
	if err != nil {
		return
	}
	err = u.matchRepo.InsertMany(ctx, matches)
	return
}

// dataCsv struct represents parsed CSV data.
type dataCsv struct {
	CompetitionID string
	Date          time.Time
	Team1         string
	Team1Score    int64
	Team2         string
	Team2Score    int64
}

// fetchDataCsv method converts CSV rows into structured data.
func (u *CompetitionUsecaseImpl) fetchDataCsv(rows []string) (data dataCsv, err error) {
	pattern := "^[a-zA-Z0-9_]+$"
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(rows[0]) {
		err = errors.Join(ErrMustRollback, fmt.Errorf("invalid format id, format must be alfanumeri alphanumeric with underscore"))
		return
	}

	timeCompetition, err := time.Parse("2006-01-02", rows[1])
	if err != nil {
		err = errors.Join(ErrCommitAndStop, fmt.Errorf("invalid time competition %s, this time example %s", rows[1], "2006-01-02"))
		return
	}

	score1, err := strconv.Atoi(rows[3])
	if err != nil {
		err = errors.Join(ErrCommitAndStop, fmt.Errorf("error converting string to integer for competition score: %v", err))
		return
	}

	score2, err := strconv.Atoi(rows[5])
	if err != nil {
		err = errors.Join(ErrCommitAndStop, fmt.Errorf("error converting string to integer for competition score: %v", err))
		return
	}

	data.CompetitionID = rows[0]
	data.Date = timeCompetition
	data.Team1 = rows[2]
	data.Team1Score = int64(score1)
	data.Team2 = rows[4]
	data.Team2Score = int64(score2)

	if len(data.Team1) > 64 || len(data.Team2) > 64 {
		err = errors.Join(ErrCommitAndStop, errors.New("error: Match names cannot exceed 64 characters"))
		return
	}

	return
}
