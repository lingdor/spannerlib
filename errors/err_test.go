package errors

import (
	"fmt"
	"github.com/lingdor/spannerlib/str"
	"testing"
)

func TestError1(t *testing.T) {

	err := fmt.Errorf("Exception MSG")
	myerr := Wrap(err, 0, "msg")
	msg := fmt.Sprintf("%v", myerr)
	if !str.Startwith(msg, "Exception MSG\ntesting.tRunner") {
		t.Error("Error format assert failed")
	}
}
