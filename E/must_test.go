package E

import (
	"strconv"
	"testing"
)

func TestMust(t *testing.T) {

	errMsg := "var1 value format failed"
	defer func() {
		err := recover()
		if dat, ok := GetErrorData[string](err); ok {
			if dat == errMsg {
				return
			}
			t.Errorf("not expect error data:%s", dat)
			return
		}
		t.Errorf("not expect error:%v", err)
	}()
	SetMustErrorData(errMsg)
	Must1(strconv.Atoi("abc"))
}
