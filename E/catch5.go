package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T any] struct {
	dataF func() DataT
	p1    P1T
	p2    P2T
	p3    P3T
	p4    P4T
	p5    P5T
	err   error
}

func (m *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T]) IfErrorDataFunc(f func() DataT) *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T] {
	m.dataF = f
	return m
}

func (m *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T]) IfErrorData(errData DataT) *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T]) Do(f func(err error)) (P1T, P2T, P3T, P4T, P5T) {
	if m.err != nil {
		f(m.err)
	}
	return m.p1, m.p2, m.p3, m.p4, m.p5
}
func (m *Catch5Box[DataT, P1T, P2T, P3T, P4T, P5T]) Must() (P1T, P2T, P3T, P4T, P5T) {
	if m.err == nil {
		return m.p1, m.p2, m.p3, m.p4, m.p5
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch5[P1T, P2T, P3T, P4T, P5T any](p1 P1T, p2 P2T, p3 P3T, p4 P4T, p5 P5T, err error) *Catch5Box[any, P1T, P2T, P3T, P4T, P5T] {

	return &Catch5Box[any, P1T, P2T, P3T, P4T, P5T]{
		err: err,
		p1:  p1,
		p2:  p2,
		p3:  p3,
		p4:  p4,
		p5:  p5,
	}
}
