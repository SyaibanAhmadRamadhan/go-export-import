package model

// Match is a DAO (data transfer object) for table match.
// this struct has methods generated from codegen https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/ginfra/gdb/generator.go.
// similar to the setter and getter features in java
type Match struct {
	ID            int64  `db:"id"`
	CompeititonID string `db:"competition_id"`
	Team1Name     string `db:"team1_name"`
	Team1Score    int64  `db:"team1_score"`
	Team2Name     string `db:"team2_name"`
	Team2Score    int64  `db:"team2_score"`
}
