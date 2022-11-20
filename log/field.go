package log

import "fmt"

type Field struct {
	Key   string
	Value any
}

func (f Field) String() string {
	return fmt.Sprintf("%s=%s", f.Key, f.Value)
}
