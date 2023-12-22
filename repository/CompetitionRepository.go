package repository

import (
	"context"
	"fmt"

	_ "github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gpostgre"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-export-import/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . CompetitionRepository
//counterfeiter:generate . MatchRepository

// CompetitionRepository is the interface for the competition table datastore
type CompetitionRepository interface {
	// InsertMany is a method of competitionRepositoryImpl that inserts multiple competition records.
	// It uses the ON CONFLICT DO NOTHING clause to handle conflicts based on the 'id' column.
	InsertMany(ctx context.Context, competitions []model.Competition) (err error)

	// LeaderBoard retrieves and displays the leaderboard for a specific competition.
	LeaderBoard(ctx context.Context, competitionID string) (res []LeaderBoardResult, err error)
}

type LeaderBoardResult struct {
	TeamName                      string
	Play, Win, Draw, Lose, Points int
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

	query, _, err := c.db.Builder.Insert(model.CompetitionTableName).
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

// LeaderBoard is an implementation of the LeaderBoard method for competitionRepositoryImpl,
// adhering to the contracts specified in the CompetitionRepository interface.
// LeaderBoard retrieves and displays the leaderboard for a specific competition.
func (c *competitionRepositoryImpl) LeaderBoard(ctx context.Context, competitionID string) (res []LeaderBoardResult, err error) {
	query := `
		SELECT team_name,
			   COALESCE(SUM(play), 0) as play,
			   COALESCE(SUM(win), 0) as win,
			   COALESCE(SUM(draw), 0) as draw,
			   COALESCE(SUM(lose), 0) as lose,
			   COALESCE(SUM(points), 0) as points
		FROM (
			SELECT team_name,
				   COUNT(*) as play,
				   SUM(CASE WHEN team_score > opponent_score THEN 1 ELSE 0 END) as win,
				   SUM(CASE WHEN team_score = opponent_score THEN 1 ELSE 0 END) as draw,
				   SUM(CASE WHEN team_score < opponent_score THEN 1 ELSE 0 END) as lose,
				   SUM(CASE WHEN team_score > opponent_score THEN 3 WHEN team_score = opponent_score THEN 1 ELSE 0 END) as points
			FROM (
				SELECT team1_name as team_name, team1_score as team_score, team2_score as opponent_score
				FROM match
				WHERE competition_id = $1
				UNION ALL
				SELECT team2_name as team_name, team2_score as team_score, team1_score as opponent_score
				FROM match
				WHERE competition_id = $1
			) as matches
			GROUP BY team_name, team_score
		) as m
		GROUP BY team_name;
	`

	rows, err := c.db.Commander.Query(ctx, query, competitionID)
	if err != nil {
		return
	}
	defer rows.Close()

	res = make([]LeaderBoardResult, 0)
	// Menampilkan hasil query
	for rows.Next() {
		var leaderBoardRes LeaderBoardResult

		err = rows.Scan(&leaderBoardRes.TeamName, &leaderBoardRes.Play, &leaderBoardRes.Win, &leaderBoardRes.Draw,
			&leaderBoardRes.Lose, &leaderBoardRes.Points)
		if err != nil {
			return
		}

		res = append(res, leaderBoardRes)
	}

	return
}
