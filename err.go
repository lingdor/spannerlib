package spannerlib

import "github.com/lingdor/spannerlib/errors"

func CheckPanic(err error) {
	if err != nil {
		panic(errors.Wrap(err, 1, "Check panic"))
	}
}
