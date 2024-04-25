package E

import (
	"github.com/lingdor/spannerlib/errors"
)

type Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T any] struct {
	dataF func() DataT
	p1    P1T
	p2    P2T
	p3    P3T
	p4    P4T
	p5    P5T
	p6    P6T
	p7    P7T
	p8    P8T
	p9    P9T
	err   error
}

func (m *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T]) IfErrorDataFunc(f func() DataT) *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T] {
	m.dataF = f
	return m
}

func (m *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T]) IfErrorData(errData DataT) *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T] {
	m.dataF = func() DataT {
		return errData
	}
	return m
}
func (m *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T]) Do(f func(err error)) (P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T) {
	if m.err != nil {
		f(m.err)
	}
	return m.p1, m.p2, m.p3, m.p4, m.p5, m.p6, m.p7, m.p8, m.p9
}
func (m *Catch9Box[DataT, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T]) Must() (P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T) {
	if m.err == nil {
		return m.p1, m.p2, m.p3, m.p4, m.p5, m.p6, m.p7, m.p8, m.p9
	}
	var errData DataT
	if m.dataF != nil {
		errData = m.dataF()
	}
	err := errors.Wrap(m.err, 1, errData)
	panic(err)
}

func Catch9[P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T any](p1 P1T, p2 P2T, p3 P3T, p4 P4T, p5 P5T, p6 P6T, p7 P7T, p8 P8T, p9 P9T, err error) *Catch9Box[any, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T] {

	return &Catch9Box[any, P1T, P2T, P3T, P4T, P5T, P6T, P7T, P8T, P9T]{
		err: err,
		p1:  p1,
		p2:  p2,
		p3:  p3,
		p4:  p4,
		p5:  p5,
		p6:  p6,
		p7:  p7,
		p8:  p8,
		p9:  p9,
	}
}
