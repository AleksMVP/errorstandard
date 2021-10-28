package errorstandard

import (
	"fmt"
	"testing"
	"time"
)

type i int32

const (
	NoError i = iota
	Expected
	Unexpected
)

// My custom expected error
var _ IError = &myerr{}

type myerr struct {
	message string
}

func (m *myerr) Error() string {
	return m.message
}

func (*myerr) Expected() bool {
	return true
}

func foo(t i) (int, IError) {
	switch t {
	case NoError:
		return 0, nil
	case Expected:
		return 0, &myerr{message: fmt.Sprintf("foo creates expected error %v", time.Now())}
	case Unexpected:
		return 0, NewUErrorF("errorstandard", "boo", nil, "foo creates unexpected error %v %v", time.Now())
	}
	return 0, nil
}

func boo(t i) (res int, err IError) {
	switch res, err = foo(t); err.(type) {
	case nil:
		break
	case *UError:
		return 0, NewUErrorF("errorstandard", "boo", err, "boo gets unexpected error from foo %v", time.Now())
	default: // expected
		return 0, err // just return untouched
	}

	return res, nil
}

func zoo() (res int, e IError) {
	_, err := boo(Unexpected)

	return 0, NewUErrorF("errorstandard", "zoo", err, "zoo gets unexpected error from boo %v", time.Now())
}

func TestNoError(t *testing.T) {
	res, err := boo(NoError)
	fmt.Printf("%v, %v\n", res, err)
}

func TestExpected(t *testing.T) {
	_, err := boo(Expected)
	fmt.Printf("%s\n", err.Error())
}

func TestUnexpected(t *testing.T) {
	_, err := zoo()
	fmt.Printf("%s\n", err.Error())
}
