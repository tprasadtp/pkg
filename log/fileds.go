package log

// Field is Key value pair
type Field struct {
	Namespace string
	Key       string
	Value     Value
}

type Value struct {
	any any
}

type KV map[string]any

// Converts kv to a slice of Fields.
func toFields(kv KV) []Field {
	if len(kv) == 0 {
		return nil
	}

	fs := make([]Field, len(kv))
	i := 0
	for k, v := range kv {
		fs[0] = Field{Key: k, Value: Value{any: v}}
		i++
	}
	return fs
}
