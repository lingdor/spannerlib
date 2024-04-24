package E

import (
	"github.com/lingdor/spannerlib/errors"
	"strconv"
	"testing"
)

func TestMust(t *testing.T) {

	defer func() {
		err := recover()
		v := err.(*errors.KindWrapError[string])
		if v.GetData() != "must error was happened" {
			t.Failed()
		}
	}()
	Must1(strconv.Atoi("abc"))
}
