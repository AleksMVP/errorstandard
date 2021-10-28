package errorstandard

type IError interface {
	Expected() bool
	error
}
