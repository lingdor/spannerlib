package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type CatchBox[DataT any] struct {
	dataF func() DataT
	err   error
}

func (m *CatchBox[DataT]) IfErrorDataFunc(f func() DataT) *CatchBox[DataT] {
	m.dataF = f
	return m
}

func (m *CatchBox[DataT]) IfErrorData(errData DataT) *CatchBox[DataT] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *CatchBox[DataT]) Do(f func(err error)) {
	if m.err != nil {
		f(m.err)
	}
}
func (m *CatchBox[DataT]) Must() {
	if m.err == nil {
		return
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch(err error) *CatchBox[any] {

	return &CatchBox[any]{
		err: err,
	}
}
