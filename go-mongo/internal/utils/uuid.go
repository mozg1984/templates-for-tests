package utils

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var UUID = reflect.TypeOf(uuid.UUID{})

func EncodeUUIDValue(_ bson.EncodeContext, vw bson.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != UUID {
		return bson.ValueEncoderError{Name: "EncodeUUIDValue", Types: []reflect.Type{UUID}, Received: val}
	}

	data := make([]byte, 16)
	for idx := 0; idx < val.Len(); idx++ {
		data[idx] = val.Index(idx).Interface().(byte)
	}

	return vw.WriteBinaryWithSubtype(data, 0x4)
}

func DecodeUUIDValue(_ bson.DecodeContext, vr bson.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != UUID {
		return bson.ValueDecoderError{Name: "DecodeUUIDValue", Types: []reflect.Type{UUID}, Received: val}
	}

	if vr.Type() != bson.TypeBinary {
		return fmt.Errorf("cannot decode %v into an UUID", vr.Type())
	}

	data, subtype, err := vr.ReadBinary()
	if err != nil {
		return err
	}

	if len(data) != 16 {
		return fmt.Errorf("DecodeUUIDValue cannot decode binary, invalid length %v", len(data))
	}

	if subtype != 0x4 {
		return fmt.Errorf("DecodeUUIDValue can only be used to decode subtype 0x4 for %s, got %v", bson.TypeBinary, subtype)
	}

	id := uuid.UUID{}
	for idx := 0; idx < 16; idx++ {
		id[idx] = data[idx]
	}
	val.Set(reflect.ValueOf(id))

	return nil
}
