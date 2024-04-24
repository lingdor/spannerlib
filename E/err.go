package E

import "github.com/lingdor/spannerlib/errors"

// GetErrorData Get Error data if err is KindWrapError
// Recommend:
//
//	defer func(){
//	  if err:=recover(); err!=nil {
//	     if dat:=E.GetErrorData(err);dat!=nil {
//
//	     }
//	  }
//	}
func GetErrorData[T any](err any) (T, bool) {
	if v, ok := err.(*errors.KindWrapError[T]); ok {
		return v.GetData(), true
	}
	if v, ok := err.(errors.ErrData); ok {
		v, vok := v.GetDataInterface().(T)
		return v, vok
	}
	var t T
	return t, false
}

// GetErrorDataInterface Get Data from KindWrapError to any
func GetErrorDataInterface(err any) (any, bool) {
	if v, ok := err.(errors.ErrData); ok {
		return v.GetDataInterface(), true
	}
	return nil, false
}
