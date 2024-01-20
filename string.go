package spannerlib

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNoFoundStartTag = errors.New("no found start tag")
var ErrNoFoundEndTag = errors.New("no found end tag")

// StringPick to cut passage of string from tags
func StringPick(str, startTag, endTag string) (string, error) {
	var index1, index2 int
	if index1 = strings.Index(str, startTag); index1 > -1 {
		index1 = index1 + len(startTag)
		if index2 = strings.Index(str[index1:], endTag); index2 > -1 {
			return str[index1 : index1+index2], nil
		} else {
			return "", fmt.Errorf("%w:%s", ErrNoFoundEndTag, endTag)
		}
	}
	return "", fmt.Errorf("%w:%s", ErrNoFoundStartTag, startTag)
}

// StringPickStart to cut passage of string to end tag
func StringPickStart(str, endTag string) (string, error) {
	var index2 int
	if index2 = strings.Index(str, endTag); index2 > -1 {
		return str[0:index2], nil
	} else {
		return "", fmt.Errorf("%w:%s", ErrNoFoundEndTag, endTag)
	}
}
