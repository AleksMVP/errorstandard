package errorstandard

type StackError struct {
	errors []ErrorWrapper
}

func (instance *StackError) PushError(err ErrorWrapper) {
	instance.errors = append(instance.errors, err)
}

func (instance *StackError) IsEmpty() bool {
	return len(instance.errors) == 0
}

func (instance *StackError) GetErrors() []ErrorWrapper {
	return instance.errors
}

func (instance *StackError) Error() (result string) {
	for _, err := range instance.errors {
		result += err.Error() + "\n"
	}

	return result
}