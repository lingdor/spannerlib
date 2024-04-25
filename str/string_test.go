package str

import (
	"testing"
)

func TestStringPick(t *testing.T) {
	body, _ := StringPick("<html><body>123</body></html>", "<body>", "</body>")
	if body != "123" {
		t.Errorf("string pick result assert failed")
	}
}
