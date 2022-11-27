package log

// KV is an alias for KV
type KV map[string]any

// Convert KV to fields slice
func (kv KV) Fields() []Field {
	return nil
}

// Field is Key value pair
type Field struct {
	Namespace string
	Key       string
	Value     Value
}

type Value struct {
	any any
}
