package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNodeClaimed = errors.New("node has been already claimed and is not in the node pool anymore")
	ErrServerSide  = errors.New("an unexpected server-side error occurred")
	ErrForbidden   = errors.New("you are not authorized to perform this operation")
	ErrNotFound    = func(entity string) error {
		return fmt.Errorf("%s not found", entity)
	}
	ErrFieldRequired = errors.New("required filed missing")
)
