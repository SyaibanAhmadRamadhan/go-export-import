package model

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

import (
	"errors"

	"time"
)

// CompetitionTableName this table or collection name
const CompetitionTableName string = "competition"

// NewCompetition is a struct with pointer that represents the table Competition in the database.
func NewCompetition() *Competition {
	return &Competition{}
}

// NewCompetitionWithOutPtr is a struct without pointer that represents the table Competition in the database.
func NewCompetitionWithOutPtr() Competition {
	return Competition{}
}

// FieldID is a field or column in the table Competition.
func (c *Competition) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table Competition.
func (c *Competition) SetID(param string) {
	c.ID = param
}

// FieldDate is a field or column in the table Competition.
func (c *Competition) FieldDate() string {
	return "date"
}

// SetDate is a setter for the field or column Date in the table Competition.
func (c *Competition) SetDate(param time.Time) {
	c.Date = param
}

// AllField is a function to get all field or column in the table Competition.
func (c *Competition) AllField() (str []string) {
	str = []string{ 
		`id`,
		`date`,
	}
	return
}

// OrderFields is a function to get all field or column in the table Competition.
func (c *Competition) OrderFields() (str []string) {
	str = []string{ 
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table Competition.
func (c *Competition) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case c.FieldID():
			values = append(values, c.ID)
		case c.FieldDate():
			values = append(values, c.Date)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct Competition.
func (c *Competition) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case c.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			c.SetID(val)
		case c.FieldDate():
			val, ok := value.(time.Time)
			if !ok {
				return errors.New("invalid type time.Time. field Date")
			}
			c.SetDate(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

