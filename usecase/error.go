package usecase

import (
	"errors"
)

// ErrMustRollback is an error type used to indicate that a transaction must be rolled back and stopped program.
var ErrMustRollback = errors.New("data will be rollback.")

// ErrCommitAndStop is an error type used to indicate that a transaction must be comit and stopped program.
var ErrCommitAndStop = errors.New("previous data wil be save.")
