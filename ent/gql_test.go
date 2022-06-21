package ent

import (
	"bytes"
	"log"
	"testing"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

func TestGqlCursor(t *testing.T) {
	cursor := UserOrderFieldBirth.toCursor(&User{ID: "abc", Birth: nil})
	if cursor.Value == nil {
		t.Fail()
	}
	cursor.MarshalGQL(bytes.NewBuffer(nil))
}

/*
type Cursor struct {
	ID    models.ID `msgpack:"i"`
	Value Value     `msgpack:"v,omitempty"`
}
*/
func TestMsgPackPanic(t *testing.T) {
	var v Value
	v = (*time.Time)(nil)
	c := &Cursor{
		Value: v,
	}
	//log.Println("OK Here")
	//_ = msgpack.NewEncoder(bytes.NewBuffer(nil)).Encode(v)
	//rv := reflect.ValueOf(c.Value)
	//rv.Kind()
	log.Println("Next panic")
	_ = msgpack.NewEncoder(bytes.NewBuffer(nil)).Encode(c)
}
