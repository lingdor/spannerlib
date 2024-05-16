package E

import (
	"github.com/lingdor/spannerlib/errors"
)

// MustTrue If return error, will panic
func MustTrue(ok bool) {
	if !ok {
		panic(errors.Wrap[any](NotOKError, 1, nil))
	}
}

// MustTrue1 If error no nil, will panic, use for 1 parameter return
// Demo:
//
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Printf("%v", err)
//		}
//	}()
//
// var num = E.MustTrue1(strconv.Atoi("123"))
func MustTrue1[T any](param T, ok bool) T {
	MustTrue(ok)
	return param
}

// MustTrue2  If error no nil, will panic, use for 2 parameter return
func MustTrue2[T, F any](param1 T, param2 F, ok bool) (T, F) {
	MustTrue(ok)
	return param1, param2
}

// MustTrue3  If error no nil, will panic, use for 3 parameter return
func MustTrue3[T, F, G any](param1 T, param2 F, param3 G, ok bool) (T, F, G) {
	MustTrue(ok)
	return param1, param2, param3
}

// MustTrue4  If error no nil, will panic, use for 4 parameter return
func MustTrue4[T, F, G, H any](param1 T, param2 F, param3 G, param4 H, ok bool) (T, F, G, H) {
	MustTrue(ok)
	return param1, param2, param3, param4
}

// MustTrue5  If error no nil, will panic, use for 5 parameter return
func MustTrue5[T, F, G, H, I any](param1 T, param2 F, param3 G, param4 H, param5 I, ok bool) (T, F, G, H, I) {
	MustTrue(ok)
	return param1, param2, param3, param4, param5
}

// MustTrue6  If error no nil, will panic, use for 6 parameter return
func MustTrue6[T, F, G, H, I, J any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, ok bool) (T, F, G, H, I, J) {
	MustTrue(ok)
	return param1, param2, param3, param4, param5, param6
}

// MustTrue7  If error no nil, will panic, use for 7 parameter return
func MustTrue7[T, F, G, H, I, J, K any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, ok bool) (T, F, G, H, I, J, K) {
	MustTrue(ok)
	return param1, param2, param3, param4, param5, param6, param7
}

// MustTrue8  If error no nil, will panic, use for 8 parameter return
func MustTrue8[T, F, G, H, I, J, K, L any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, param8 L, ok bool) (T, F, G, H, I, J, K, L) {
	MustTrue(ok)
	return param1, param2, param3, param4, param5, param6, param7, param8
}

// MustTrue9  If error no nil, will panic, use for 9 parameter return
func MustTrue9[T, F, G, H, I, J, K, L, M any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, param8 L, param9 M, ok bool) (T, F, G, H, I, J, K, L, M) {
	MustTrue(ok)
	return param1, param2, param3, param4, param5, param6, param7, param8, param9
}
