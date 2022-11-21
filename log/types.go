package log

type StackTrace struct{}

type CallerInfo struct {
	Package string
	File    string
	Line    string
}

type Field struct {
	Namespace string
	Key       string
	Value     any
}
