package errorstandard

import "fmt"

type ErrorWrapper struct {
	Original error
	Type     error
}

func NewErrorWrapper(original error, eType error) ErrorWrapper {
	return ErrorWrapper{
		Original: original,
		Type:     eType,
	}
}

func (instance *ErrorWrapper) Error() string {
	return fmt.Sprintf(
		"Original: %v, Type: %v",
		instance.Original.Error(),
		instance.Type.Error(),
	)
}
