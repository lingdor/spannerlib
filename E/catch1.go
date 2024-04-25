package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch1Box[DataT, P1T any] struct {
	dataF func() DataT
	p1    P1T
	err   error
}

func (m *Catch1Box[DataT, P1T]) IfErrorDataFunc(f func() DataT) *Catch1Box[DataT, P1T] {
	m.dataF = f
	return m
}

func (m *Catch1Box[DataT, P1T]) IfErrorData(errData DataT) *Catch1Box[DataT, P1T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch1Box[DataT, P1T]) Do(f func(err error)) P1T {
	if m.err != nil {
		f(m.err)
	}
	return m.p1
}
func (m *Catch1Box[DataT, P1T]) Must() P1T {
	if m.err == nil {
		return m.p1
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch1[P1T any](p1 P1T, err error) *Catch1Box[any, P1T] {

	return &Catch1Box[any, P1T]{
		err: err,
		p1:  p1,
	}
}
