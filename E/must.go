package E

import (
	"github.com/lingdor/spannerlib/errors"
)

// Must If return error, will panic
func Must(err error) {
	if err != nil {
		panic(errors.Wrap(err, 2, "must error was happened"))
	}

}

// Must1 If error no nil, will panic, use for 1 parameter return
// Demo:
//
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Printf("%v", err)
//		}
//	}()
//
// var num = E.Must1(strconv.Atoi("123"))
func Must1[T any](param T, err error) T {
	Must(err)
	return param
}

// Must2  If error no nil, will panic, use for 2 parameter return
func Must2[T, F any](param1 T, param2 F, err error) (T, F) {
	Must(err)
	return param1, param2
}

// Must3  If error no nil, will panic, use for 3 parameter return
func Must3[T, F, G any](param1 T, param2 F, param3 G, err error) (T, F, G) {
	Must(err)
	return param1, param2, param3
}

// Must4  If error no nil, will panic, use for 4 parameter return
func Must4[T, F, G, H any](param1 T, param2 F, param3 G, param4 H, err error) (T, F, G, H) {
	Must(err)
	return param1, param2, param3, param4
}

// Must5  If error no nil, will panic, use for 5 parameter return
func Must5[T, F, G, H, I any](param1 T, param2 F, param3 G, param4 H, param5 I, err error) (T, F, G, H, I) {
	Must(err)
	return param1, param2, param3, param4, param5
}

// Must6  If error no nil, will panic, use for 6 parameter return
func Must6[T, F, G, H, I, J any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, err error) (T, F, G, H, I, J) {
	Must(err)
	return param1, param2, param3, param4, param5, param6
}

// Must7  If error no nil, will panic, use for 7 parameter return
func Must7[T, F, G, H, I, J, K any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, err error) (T, F, G, H, I, J, K) {
	Must(err)
	return param1, param2, param3, param4, param5, param6, param7
}

// Must8  If error no nil, will panic, use for 8 parameter return
func Must8[T, F, G, H, I, J, K, L any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, param8 L, err error) (T, F, G, H, I, J, K, L) {
	Must(err)
	return param1, param2, param3, param4, param5, param6, param7, param8
}

// Must9  If error no nil, will panic, use for 9 parameter return
func Must9[T, F, G, H, I, J, K, L, M any](param1 T, param2 F, param3 G, param4 H, param5 I, param6 J, param7 K, param8 L, param9 M, err error) (T, F, G, H, I, J, K, L, M) {
	Must(err)
	return param1, param2, param3, param4, param5, param6, param7, param8, param9
}
