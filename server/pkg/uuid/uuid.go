package uuid

import (
	"fmt"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// This is a value (de|en)coder for the github.com/google/uuid UUID type. For best experience, register
// mongoRegistry to mongo client instance via options, e.g.
//  clientOptions := options.Client().SetRegistry(mongoRegistry)
//
// Only BSON binary subtype 0x04 is supported.
//
// Use as you please

var (
	tUUID       = reflect.TypeOf(uuid.UUID{})
	uuidSubtype = byte(0x04)

	MongoRegistry = bson.NewRegistryBuilder().
			RegisterTypeEncoder(tUUID, bsoncodec.ValueEncoderFunc(uuidEncodeValue)).
			RegisterTypeDecoder(tUUID, bsoncodec.ValueDecoderFunc(uuidDecodeValue)).
			Build()
)

func uuidEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tUUID {
		return bsoncodec.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}
	b := val.Interface().(uuid.UUID)
	return vw.WriteBinaryWithSubtype(b[:], uuidSubtype)
}

func uuidDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tUUID {
		return bsoncodec.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}

	var data []byte
	var subtype byte
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.Binary:
		data, subtype, err = vr.ReadBinary()
		if subtype != uuidSubtype {
			return fmt.Errorf("unsupported binary subtype %v for UUID", subtype)
		}
	case bsontype.Null:
		err = vr.ReadNull()
	case bsontype.Undefined:
		err = vr.ReadUndefined()
	default:
		return fmt.Errorf("cannot decode %v into a UUID", vrType)
	}

	if err != nil {
		return err
	}
	uuid2, err := uuid.UUIDFromBytes(data)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(uuid2))
	return nil
}
