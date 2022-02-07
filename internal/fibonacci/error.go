package fibonacci

import "errors"

var (
	ErrBadIndex = errors.New("the 'from' is greater than the 'to'")
)
