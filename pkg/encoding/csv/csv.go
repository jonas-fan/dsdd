package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"reflect"
)

type Assigner interface {
	Assign(keys []string, values []string)
}

var AssignerType = reflect.TypeOf((*Assigner)(nil)).Elem()

func assignerType(pointer interface{}) (reflect.Type, error) {
	var t reflect.Type

	if t = reflect.TypeOf(pointer); t.Kind() != reflect.Ptr {
		return nil, errors.New("Not a pointer")
	}

	if t = t.Elem(); t.Kind() != reflect.Slice {
		return nil, errors.New("Not a pointer to slice")
	}

	if t = t.Elem(); !reflect.PtrTo(t).Implements(AssignerType) {
		return nil, errors.New("AssignerType not implemented")
	}

	return t, nil
}

func readAll(reader *csv.Reader, pointer interface{}) error {
	elementType, err := assignerType(pointer)

	if err != nil {
		return err
	}

	var out = reflect.ValueOf(pointer).Elem()
	var keys []string

	for {
		values, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if keys == nil {
			keys = values
			continue
		}

		ptr := reflect.New(elementType)
		ptr.Interface().(Assigner).Assign(keys, values)
		out.Set(reflect.Append(out, ptr.Elem()))
	}

	return nil
}

// ReadAll reads a whole CSV file then fills a pointer to slice of objects
// that implement Assigner interface.
func ReadAll(reader io.Reader, pointer interface{}) error {
	return readAll(csv.NewReader(reader), pointer)
}
