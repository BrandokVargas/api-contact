package validErrors

import "errors"

var (
	ErrPersonCanNotBeNil = errors.New("the user can't be null")
)
