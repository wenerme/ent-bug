package ent

import (
	"bytes"
	"testing"
)

func TestGqlCursor(t *testing.T) {
	UserOrderFieldBirth.toCursor(&User{ID: "abc", Birth: nil}).MarshalGQL(bytes.NewBuffer(nil))
}
