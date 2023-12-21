package model

import (
	"time"
)

// Competition is a DAO (data transfer object) for table competition.
// this struct has methods generated from codegen https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/ginfra/gdb/generator.go.
// similar to the setter and getter features in java
type Competition struct {
	ID   string    `db:"id"`
	Date time.Time `db:"date"`
}
