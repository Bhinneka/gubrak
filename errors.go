package gubrak

import (
	"errors"
)

var (
	// ErrorConfigNotFound for error config file not found
	ErrorConfigNotFound = errors.New("config file not found")

	// ErrorFormatConfigNotValid for invalid format config error
	ErrorFormatConfigNotValid = errors.New("format config is not valid")
)
