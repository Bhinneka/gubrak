package gubrak

import (
	"errors"
)

var (
	ErrorConfigNotFound       = errors.New("config file not found")
	ErrorFormatConfigNotValid = errors.New("format config is not valid")
)
