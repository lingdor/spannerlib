package errors

import (
	"bytes"
	"fmt"
	"github.com/lingdor/spannerlib/str"
	"runtime"
)

const MaxStackDepth = 50
const pkgName = "github.com/lingdor/spannerlib"

type KindWrapError[T any] struct {
	kind  error
	data  T
	stack []uintptr
}

func (k *KindWrapError[T]) Error() string {
	return k.kind.Error()
}
func (k *KindWrapError[T]) Format(f fmt.State, verb rune) {
	if verb == 'v' {
		_, _ = f.Write([]byte(k.ErrorStack()))
	} else {
		_, _ = f.Write([]byte(k.Error()))
	}
}
func (k *KindWrapError[T]) GetData() T {
	return k.data
}

func (k *KindWrapError[T]) GetDataInterface() any {
	return k.data
}

func (k *KindWrapError[T]) ErrorStack() string {
	return k.Error() + "\n" + string(k.Stack())
}
func (k *KindWrapError[T]) Stack() []byte {
	buf := bytes.Buffer{}
	frames := runtime.CallersFrames(k.stack)
	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		if str.Startwith(frame.Function, pkgName) {
			continue
		}
		buf.WriteString(fmt.Sprintf("%s(%s:%d)\n", frame.Function, frame.File, frame.Line))
	}
	return buf.Bytes()
}

func Wrap[T any](kind error, skip int, data T) *KindWrapError[T] {

	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2+skip, stack[:])
	return &KindWrapError[T]{
		kind:  kind,
		data:  data,
		stack: stack[:length],
	}
}
