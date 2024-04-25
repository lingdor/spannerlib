package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch4Box[DataT, P1T, P2T, P3T, P4T any] struct {
	dataF func() DataT
	p1    P1T
	p2    P2T
	p3    P3T
	p4    P4T
	err   error
}

func (m *Catch4Box[DataT, P1T, P2T, P3T, P4T]) IfErrorDataFunc(f func() DataT) *Catch4Box[DataT, P1T, P2T, P3T, P4T] {
	m.dataF = f
	return m
}

func (m *Catch4Box[DataT, P1T, P2T, P3T, P4T]) IfErrorData(errData DataT) *Catch4Box[DataT, P1T, P2T, P3T, P4T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch4Box[DataT, P1T, P2T, P3T, P4T]) Do(f func(err error)) (P1T, P2T, P3T, P4T) {
	if m.err != nil {
		f(m.err)
	}
	return m.p1, m.p2, m.p3, m.p4
}
func (m *Catch4Box[DataT, P1T, P2T, P3T, P4T]) Must() (P1T, P2T, P3T, P4T) {
	if m.err == nil {
		return m.p1, m.p2, m.p3, m.p4
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch4[P1T, P2T, P3T, P4T any](p1 P1T, p2 P2T, p3 P3T, p4 P4T, err error) *Catch4Box[any, P1T, P2T, P3T, P4T] {

	return &Catch4Box[any, P1T, P2T, P3T, P4T]{
		err: err,
		p1:  p1,
		p2:  p2,
		p3:  p3,
		p4:  p4,
	}
}
