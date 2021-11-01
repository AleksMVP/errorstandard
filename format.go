package errorstandard

import "fmt"

func FormatError(methodName, meta string, err error) error {
	return fmt.Errorf(
		"error in %v, with meta: %v, err: %v",
		methodName,
		meta,
		err,
	)
}
