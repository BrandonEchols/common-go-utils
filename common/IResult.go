package common

/*
	IResult is the wrapper for a struct that is meant to be returned from a function as a more verbose error.
*/
type IResult interface {
	GetChild() IResult
	GetChildren() []chan asyncLogPackage
	WasSuccessful() bool
	Succeed()
	Fail()
	Error() string
	MergeWithResult(r IResult)
	GetMessages() []string
	GetLogLevel() int
	GetStatusCode() int
	SetStatusCode(int)
	GetResponseMessage() string
	SetResponseMessage(string)
	Flush()
	Debugf(template string, args ...interface{})
	DebugMessagef(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}
