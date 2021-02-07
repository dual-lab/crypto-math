package err

import "errors"

var (
	ErrInvalidArg       = errors.New("invalid argument")
	ErrInvaliTypedArg   = errors.New("invalid argument type")
	ErrInvalidFlagValue = errors.New("Invalid flag value")
)
