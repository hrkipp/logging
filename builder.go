package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type Builder func(ctx context.Context, level Level, args ...interface{}) []byte

func DefaultBuilder(ctx context.Context, level Level, args ...interface{}) []byte {
	var buf bytes.Buffer
	buf.WriteString(level.String())
	switch arg := args[0].(type) {
	case string:
		str := fmt.Sprintf(arg, args[1:]...)
		buf.WriteString(" ")
		buf.WriteString(str)
	default:
		for _, arg := range args {
			buf.WriteString(" ")
			if isPrimitive(arg) {
				buf.WriteString(fmt.Sprint(arg))
			} else {
				json.NewEncoder(&buf).Encode(arg)
				buf.Truncate(buf.Len() - 1)
			}
		}
	}
	return buf.Bytes()
}

func isPrimitive(val interface{}) bool {
	switch val.(type) {
	case string, []byte,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		complex64, complex128,
		float32, float64:
		return true
	}
	return false

}
