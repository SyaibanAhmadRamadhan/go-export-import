package repository

import (
	"context"
	"fmt"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gpostgre"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-export-import/model"
)

// MatchRepository is the interface for the match table datastore.
type MatchRepository interface {
	// InsertMany is a method of MatchRepository that inserts multiple match records.
	// It receives a slice of model.Match objects and inserts them into the datastore.
	InsertMany(ctx context.Context, matches []model.Match) (err error)
}

// matchRepositoryImpl is an implementation of the MatchRepository interface.
type matchRepositoryImpl struct {
	db *gpostgre.PostgresPgx
}

// NewMatchRepositoryImpl creates a new instance of MatchRepository implementation.
// It takes a *gpostgre.PostgresPgx instance as a parameter.
// The returned instance adheres to the MatchRepository interface.
// this return matchRepositoryImpl
func NewMatchRepositoryImpl(db *gpostgre.PostgresPgx) MatchRepository {
	return &matchRepositoryImpl{
		db: db,
	}
}

// InsertMany is an implementation of the InsertMany method for matchRepositoryImpl,
// adhering to the contracts specified in the MatchRepository interface.
// It inserts multiple match records into the datastore using the specified table name (match).
// This method uses SendBatch to bulk insert into postgres.
// The method returns an error if the insertion process encounters any issues.
func (s *matchRepositoryImpl) InsertMany(ctx context.Context, matches []model.Match) (err error) {
	if matches == nil || len(matches) == 0 {
		return
	}

	matchOne := matches[0]
	var columns []string
	for _, v := range matchOne.AllField() {
		if v == "id" {
			continue
		}
		columns = append(columns, v)
	}

	emptyValues := make([]interface{}, len(columns))
	for i := range emptyValues {
		emptyValues[i] = ""
	}
	query, _, err := s.db.Builder.Insert(model.MatchTableName).Columns(columns...).
		Values(emptyValues...).ToSql()
	if err != nil {
		log.Info().Msgf("error ToSql builder, %v", err)
		return
	}

	batch := &pgx.Batch{}
	for _, match := range matches {
		values := match.GetValuesByColums(columns...)
		batch.Queue(query, values...)
	}

	result := s.db.Commander.SendBatch(ctx, batch)
	defer func() {
		if err := result.Close(); err != nil {
			log.Info().Msgf("error SendBatch close, %v", err)
		}
	}()

	for range matches {
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
