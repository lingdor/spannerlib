package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch3Box[DataT, P1T, P2T, P3T any] struct {
	dataF func() DataT
	p1    P1T
	p2    P2T
	p3    P3T
	err   error
}

func (m *Catch3Box[DataT, P1T, P2T, P3T]) IfErrorDataFunc(f func() DataT) *Catch3Box[DataT, P1T, P2T, P3T] {
	m.dataF = f
	return m
}

func (m *Catch3Box[DataT, P1T, P2T, P3T]) IfErrorData(errData DataT) *Catch3Box[DataT, P1T, P2T, P3T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch3Box[DataT, P1T, P2T, P3T]) Do(f func(err error)) (P1T, P2T, P3T) {
	if m.err != nil {
		f(m.err)
	}
	return m.p1, m.p2, m.p3
}
func (m *Catch3Box[DataT, P1T, P2T, P3T]) Must() (P1T, P2T, P3T) {
	if m.err == nil {
		return m.p1, m.p2, m.p3
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch3[P1T, P2T, P3T any](p1 P1T, p2 P2T, p3 P3T, err error) *Catch3Box[any, P1T, P2T, P3T] {

	return &Catch3Box[any, P1T, P2T, P3T]{
		err: err,
		p1:  p1,
		p2:  p2,
		p3:  p3,
	}
}
