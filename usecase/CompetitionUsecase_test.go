package usecase

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gdir"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gdbfakes"
	"github.com/stretchr/testify/assert"

	"github.com/SyaibanAhmadRamadhan/go-export-import/repository"
	"github.com/SyaibanAhmadRamadhan/go-export-import/repository/repositoryfakes"
)

func TestCompetitionUsecaseImpl_Input(t *testing.T) {
	compeRepoMock := &repositoryfakes.FakeCompetitionRepository{}
	matchRepoMock := &repositoryfakes.FakeMatchRepository{}
	txMock := &gdbfakes.FakeTx{}

	compeUsecase := NewCompetitionUsecaseImpl(compeRepoMock, matchRepoMock, txMock)

	testTables := []struct {
		dataCsv                [][]string
		compeAndMatchRepoError int
		err                    bool
		mustRollback           bool
	}{
		{
			dataCsv: [][]string{
				{"competition_id", "date", "team_1", "team_1_score", "team_2", "team_2_score"},
				{"1", "2023-11-11", "Team A", "1", "Team B", "0"},
				{"2", "2023-11-11", "Team D", "1", "Team E", "0"},
				{"3", "2023-11-11", "Team G", "1", "Team F", "0"},
				{"4", "2023-11-11", "Team X", "1", "Team Y", "0"},
				{"5", "2023-11-11", "Team Z", "1", "Team W", "0"},
				{"6", "2023-11-11", "Team M", "1", "Team N", "0"},
				{"7", "2023-11-11", "Team Q", "1", "Team R", "0"},
				{"8", "2023-11-11", "Team S", "1", "Team T", "0"},
				{"9", "2023-11-11", "Team U", "1", "Team V", "0"},
				{"10", "2023-11-11", "Team K", "1", "Team L", "0"},
			},
			compeAndMatchRepoError: 4,
			err:                    false,
			mustRollback:           false,
		},
		{
			dataCsv: [][]string{
				{"competition_id", "date", "team_1", "team_1_score", "team_2", "team_2_score"},
				{"1@", "2023-11-11", "Team A", "1", "Team B", "0"},
				{"2", "2023-11-11", "Team D", "1", "Team E", "0"},
				{"3", "2023-11-11", "Team G", "1", "Team F", "0"},
				{"4", "2023-11-11", "Team X", "1", "Team Y", "0"},
				{"5", "2023-11-11", "Team Z", "1", "Team W", "0"},
				{"6", "2023-11-11", "Team M", "1", "Team N", "s"},
				{"7", "2023-11-11", "Team Q", "1", "Team R", "0"},
				{"8", "2023-11-11", "Team S", "1", "Team T", "0"},
				{"9", "2023-11-11", "Team U", "1", "Team V", "0"},
				{"10", "2023-11-11", "Team K", "1", "Team L", "0"},
			},
			compeAndMatchRepoError: 0,
			err:                    true,
			mustRollback:           true,
		},
		{
			dataCsv: [][]string{
				{"competition_id", "date", "team_1", "team_1_score", "team_2", "team_2_score"},
				{"1", "2023-11-11", "Team A", "1", "Team B", "0"},
				{"2", "2023-11-11", "Team D", "1", "Team E", "0"},
				{"3", "2023-11-11", "Team G", "1", "Team F", "0"},
				{"4", "2023-11-11", "Team X", "1", "Team Y", "0"},
				{"5", "2023-11-11", "Team Z", "1", "Team W", "0"},
				{"6", "2023-11-11", "Team M", "1", "Team N", "sasd"},
				{"7", "2023-11-11", "Team Q", "1", "Team R", "0"},
				{"8", "2023-11-11", "Team S", "1", "Team T", "0"},
				{"9", "2023-11-11", "Team U", "1", "Team V", "0"},
				{"10", "2023-11-11", "Team K", "1", "Team L", "0"},
			},
			compeAndMatchRepoError: 0,
			err:                    true,
			mustRollback:           false,
		},
	}

	for i, table := range testTables {
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			csvFileName := "test_input" + strconv.Itoa(i) + ".csv"
			err := writeToCSV(table.dataCsv, csvFileName)
			assert.NoError(t, err, "Error writing CSV file")

			txMock.DoTransactionStub = func(ctx context.Context, option *gdb.TxOption, f func(c context.Context) (commit bool, err error)) error {
				if table.err {
					for i, data := range table.dataCsv {
						if i == 0 {
							continue
						}
						_, err := compeUsecase.fetchDataCsv(data)
						if err != nil {
							return err
						}
					}
				}
				return nil
			}

			err = compeUsecase.Input([]string{csvFileName}, 3)
			if table.err {
				if table.mustRollback {
					assert.ErrorIs(t, err, ErrMustRollback)
				} else {
					assert.ErrorIs(t, err, ErrCommitAndStop)
				}
			} else {
				assert.NoError(t, err, "Expected an no error")
			}

			assert.Equal(t, i+1, txMock.DoTransactionCallCount(), "Expected DoTransaction to be called once")
			os.Remove(csvFileName)
		})
	}
}

func TestLeaderBoard(t *testing.T) {
	// Mock data and dependencies
	compeRepoMock := &repositoryfakes.FakeCompetitionRepository{}
	matchRepoMock := &repositoryfakes.FakeMatchRepository{}
	txMock := &gdbfakes.FakeTx{}

	compeUsecase := NewCompetitionUsecaseImpl(compeRepoMock, matchRepoMock, txMock)

	err := compeUsecase.LeaderBoard(context.Background(), []string{})
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid arguments")

	compeRepoMock.LeaderBoardReturns(nil, errors.New("repository error"))
	err = compeUsecase.LeaderBoard(context.Background(), []string{"3", "filenamecompetition"})
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")

	compeRepoMock.LeaderBoardReturns([]repository.LeaderBoardResult{}, nil)
	err = compeUsecase.LeaderBoard(context.Background(), []string{"3", "filenamecompetition"})
	assert.NoError(t, err)

	dir, err := gdir.FindDirPathOfFileFromGoMod("filenamecompetition.csv")
	assert.NoError(t, err)
	assert.NotEqual(t, "", dir)

	dir, err = gdir.LocateGoModDirectory()
	assert.NoError(t, err)
	os.Remove(dir + "/res/filenamecompetition.csv")
}

func writeToCSV(dataCsv [][]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range dataCsv {
		_, err := file.WriteString(strings.Join(row, ",") + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
