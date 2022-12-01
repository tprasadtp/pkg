package log

import "strings"

// Returns a new Field. Namespace is optional,
// if multiple namespaces are specified, it is joined with a '.'
// as the separator.
func F(key string, value any, namespace ...string) Field {
	switch len(namespace) {
	case 0:
		return Field{
			Key:   key,
			Value: NewValue(value),
		}
	case 1:
		return Field{
			Namespace: namespace[0],
			Key:       key,
			Value:     NewValue(value),
		}
	default:
		return Field{
			Namespace: strings.Join(namespace, "."),
			Key:       key,
			Value:     NewValue(value),
		}
	}
}

// Field is Key value pair.
type Field struct {
	Namespace string
	Key       string
	Value     Value
}
