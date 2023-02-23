package errors

import (
	"fmt"
)

// Newf is like New except it formats the message with a fmt
// format specifier.
func Newf(format string, a ...interface{}) error {
	err := New(fmt.Sprintf(format, a...))
	SetLocation(err, 1)
	return err
}

// Notef is like Note except it formats the message with a fmt
// format specifier.
func Notef(err error, shouldPreserveCause func(error) bool, format string, a ...interface{}) error {
	err = Note(err, shouldPreserveCause, fmt.Sprintf(format, a...))
	SetLocation(err, 1)
	return err
}

// Becausef is like Because except it formats the message with a fmt
// format specifier.
func Becausef(err, cause error, format string, a ...interface{}) error {
	err = Because(err, cause, fmt.Sprintf(format, a...))
	SetLocation(err, 1)
	return err
}
