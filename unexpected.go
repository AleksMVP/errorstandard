package errorstandard

import (
	"fmt"
	"strings"
)

var _ error = &UError{}
var _ IError = &UError{}

type UError struct {
	source     string
	method     string
	format     string
	args       []interface{}
	InnerError error
	length     int32
}

func countPlaceHolders(format string) int {
	return strings.Count(format, "%") - strings.Count(format, "%%")
}

func NewUError(source string, method string, inner error, message string) *UError {
	return NewUErrorF(source, method, inner, message)
}

func NewUErrorF(source string, method string, inner error, format string, args ...interface{}) *UError {
	count := countPlaceHolders(format)
	fullArgs := make([]interface{}, count+2)
	if len(args) > count {
		args = args[:count]
	}
	_ = copy(fullArgs, args)
	fullArgs[count] = method
	fullArgs[count+1] = source
	format += "\t at %q from %q"

	if err, ok := inner.(*UError); ok {
		return &UError{
			source:     source,
			method:     method,
			format:     format,
			args:       fullArgs,
			length:     err.length + 1,
			InnerError: err,
		}
	} else {
		var length int32 = 1
		if err != nil {
			length = 2
		}
		return &UError{
			source:     source,
			method:     method,
			format:     format,
			args:       fullArgs,
			length:     length,
			InnerError: inner,
		}
	}

}

func (err *UError) Error() (result string) {
	formats := make([]string, 0, err.length)
	args := make([]([]interface{}), 0, err.length)
	argsLength := 0
	var child error
	var ok bool

	formats = append(formats, err.format)
	args = append(args, err.args)
	argsLength += len(err.args)
	child = err.InnerError
	for child != nil {
		if err, ok = child.(*UError); ok {
			formats = append(formats, err.format)
			args = append(args, err.args)
			argsLength += len(err.args)
			child = err.InnerError
		} else {
			formats = append(formats, child.Error())
			child = nil
		}
	}

	allArgs, offset := make([]interface{}, argsLength), 0
	for _, a := range args {
		offset += copy(allArgs[offset:], a)
	}

	return fmt.Sprintf(strings.Join(formats, "\n\t"), allArgs...)
}

func (err *UError) Expected() bool {
	return false
}
