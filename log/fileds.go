package log

// KV is an alias for KV
type KV map[string]any

// Convert KV to fields slice
func (kv KV) Fields() []Field {
	return nil
}

// Returns a new Field. Namespace is optional,
// if multiple namespaces are specified, it is joined with '.'
// func F(key string, value any, ns ...string) Field {
// 	if len(ns) == 0 {
//         return Field{
//             Key: ,
//         }
// 	}
// }

// Field is Key value pair
type Field struct {
	Namespace string
	Key       string
	Value     Value
}

type Value struct {
	any any
}
