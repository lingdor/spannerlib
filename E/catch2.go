package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch2Box[DataT, P1T, P2T any] struct {
	dataF func() DataT
	p1    P1T
	p2    P2T
	err   error
}

func (m *Catch2Box[DataT, P1T, P2T]) IfErrorDataFunc(f func() DataT) *Catch2Box[DataT, P1T, P2T] {
	m.dataF = f
	return m
}

func (m *Catch2Box[DataT, P1T, P2T]) IfErrorData(errData DataT) *Catch2Box[DataT, P1T, P2T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch2Box[DataT, P1T, P2T]) Do(f func(err error)) (P1T, P2T) {
	if m.err != nil {
		f(m.err)
	}
	return m.p1, m.p2
}
func (m *Catch2Box[DataT, P1T, P2T]) Must() (P1T, P2T) {
	if m.err == nil {
		return m.p1, m.p2
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch2[P1T, P2T any](p1 P1T, p2 P2T, err error) *Catch2Box[any, P1T, P2T] {

	return &Catch2Box[any, P1T, P2T]{
		err: err,
		p1:  p1,
		p2:  p2,
	}
}
