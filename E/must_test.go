package E

import (
	"strconv"
	"testing"
)

func TestMust(t *testing.T) {

	errMsg := "must method panic a error"
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
	Must1(strconv.Atoi("abc"))
}

func TestWhenError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if errrData, ok := GetErrorData[string](err); ok {
				if errrData == "custome error" {
					return
				}
			}
		}
		t.Error("recover not match")
	}()
	Catch1(strconv.Atoi("abc")).IfErrorData("custome error").Must()
}
