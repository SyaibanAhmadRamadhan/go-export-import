package repository

import (
	"context"
	"fmt"

	_ "github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gpostgre"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-export-import-big-data/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . CompetitionRepository
//counterfeiter:generate . MatchRepository

// CompetitionRepository is the interface for the competition table datastore
type CompetitionRepository interface {
	// InsertMany is a method of competitionRepositoryImpl that inserts multiple competition records.
	// It uses the ON CONFLICT DO NOTHING clause to handle conflicts based on the 'id' column.
	InsertMany(ctx context.Context, competitions []model.Competition) (err error)
}

// competitionRepositoryImpl is an inheritance of CompetitionRepository.
// that follows the contracts in the CompetitionRepository interface.
type competitionRepositoryImpl struct {
	db *gpostgre.PostgresPgx
}

// NewCompetitionRepositoryImpl creates a new instance of CompetitionRepository
// implementation that inheritance from CompetitionRepository interface.
// It ensures adherence to the contracts specified in the CompetitionRepository interface.
// It takes a *gpostgre.PostgresPgx instance as a parameter.
// this return competitionRepositoryImpl
func NewCompetitionRepositoryImpl(db *gpostgre.PostgresPgx) CompetitionRepository {
	return &competitionRepositoryImpl{
		db: db,
	}
}

// InsertMany is an implementation of the InsertMany method for competitionRepositoryImpl,
// adhering to the contracts specified in the CompetitionRepository interface.
// It inserts multiple competition records into the datastore using the ON CONFLICT DO NOTHING clause
// to handle conflicts based on the 'id' column.
// This method uses SendBatch to bulk insert into postgres.
// The method returns an error if the insertion process encounters any issues.
func (c *competitionRepositoryImpl) InsertMany(ctx context.Context, competitions []model.Competition) (err error) {
	if competitions == nil || len(competitions) == 0 {
		return
	}

	query, _, err := c.db.Builder.Insert(model.CompeititonTableName).
		Columns(competitions[0].FieldID(), competitions[0].FieldDate()).
		Suffix("ON CONFLICT(id) DO NOTHING").
		Values("", "").ToSql()
	if err != nil {
		log.Info().Msgf("error ToSql builder, %v", err)
		return
	}

	batch := &pgx.Batch{}
	for _, competition := range competitions {
		batch.Queue(query, competition.ID, competition.Date)
	}

	result := c.db.Commander.SendBatch(ctx, batch)
	defer func() {
		if err := result.Close(); err != nil {
			log.Info().Msgf("error SendBatch close, %v", err)
		}
	}()

	for range competitions {
		_, err = result.Exec()
		if err != nil {
			// var pgErr *pgconn.PgError
			// if errors.As(err, &pgErr) && pgErr.Code == pq.ErrorCode.Name() {
			// 	log.Printf("competition %s already exists", competition.ID)
			// 	continue
			// }

			return fmt.Errorf("unable to insert row: %w", err)
		}
	}

	return
}
