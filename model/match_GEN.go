package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"
)

// MatchTableName this table or collection name
const MatchTableName string = "match"

// NewMatch is a struct with pointer that represents the table Match in the database.
func NewMatch() *Match {
	return &Match{}
}

// NewMatchWithOutPtr is a struct without pointer that represents the table Match in the database.
func NewMatchWithOutPtr() Match {
	return Match{}
}

// FieldCompeititonID is a field or column in the table Match.
func (m *Match) FieldCompeititonID() string {
	return "competition_id"
}

// SetCompeititonID is a setter for the field or column CompeititonID in the table Match.
func (m *Match) SetCompeititonID(param string) {
	m.CompeititonID = param
}

// FieldTeam1Name is a field or column in the table Match.
func (m *Match) FieldTeam1Name() string {
	return "team1_name"
}

// SetTeam1Name is a setter for the field or column Team1Name in the table Match.
func (m *Match) SetTeam1Name(param string) {
	m.Team1Name = param
}

// FieldTeam1Score is a field or column in the table Match.
func (m *Match) FieldTeam1Score() string {
	return "team1_score"
}

// SetTeam1Score is a setter for the field or column Team1Score in the table Match.
func (m *Match) SetTeam1Score(param int64) {
	m.Team1Score = param
}

// FieldTeam2Name is a field or column in the table Match.
func (m *Match) FieldTeam2Name() string {
	return "team2_name"
}

// SetTeam2Name is a setter for the field or column Team2Name in the table Match.
func (m *Match) SetTeam2Name(param string) {
	m.Team2Name = param
}

// FieldTeam2Score is a field or column in the table Match.
func (m *Match) FieldTeam2Score() string {
	return "team2_score"
}

// SetTeam2Score is a setter for the field or column Team2Score in the table Match.
func (m *Match) SetTeam2Score(param int64) {
	m.Team2Score = param
}

// FieldID is a field or column in the table Match.
func (m *Match) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table Match.
func (m *Match) SetID(param int64) {
	m.ID = param
}

// AllField is a function to get all field or column in the table Match.
func (m *Match) AllField() (str []string) {
	str = []string{ 
		`team1_score`,
		`team2_name`,
		`team2_score`,
		`id`,
		`competition_id`,
		`team1_name`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Match.
func (m *Match) OrderFields() (str []string) {
	str = []string{ 
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Match.
func (m *Match) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case m.FieldTeam1Score():
			values = append(values, m.Team1Score)
		case m.FieldTeam2Name():
			values = append(values, m.Team2Name)
		case m.FieldTeam2Score():
			values = append(values, m.Team2Score)
		case m.FieldID():
			values = append(values, m.ID)
		case m.FieldCompeititonID():
			values = append(values, m.CompeititonID)
		case m.FieldTeam1Name():
			values = append(values, m.Team1Name)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Match.
func (m *Match) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case m.FieldTeam1Score():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field Team1Score")
			}
			m.SetTeam1Score(val)
		case m.FieldTeam2Name():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Team2Name")
			}
			m.SetTeam2Name(val)
		case m.FieldTeam2Score():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field Team2Score")
			}
			m.SetTeam2Score(val)
		case m.FieldID():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field ID")
			}
			m.SetID(val)
		case m.FieldCompeititonID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CompeititonID")
			}
			m.SetCompeititonID(val)
		case m.FieldTeam1Name():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Team1Name")
			}
			m.SetTeam1Name(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

