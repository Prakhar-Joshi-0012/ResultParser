package errors

import (
	"errors"
	"fmt"
)

func ErrorEncountered(err error, msg ...string) error {
	e := fmt.Sprintf("Error: %v\n", err)
	for _, msg := range msg {
		e += fmt.Sprintf("%v\n", msg)
	}
	return errors.New(e)
}
