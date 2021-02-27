package csv

import (
	"strings"
	"testing"
)

const data = `
Id,Name,Birthday
0,Linux,1991
1,Apple,1995
`

type Event struct {
	Id       string
	Name     string
	Birthday string
}

func (e *Event) Assign(keys []string, values []string) {
	for index := range keys {
		switch keys[index] {
		case "Id":
			e.Id = values[index]
		case "Name":
			e.Name = values[index]
		case "Birthday":
			e.Birthday = values[index]
		}
	}
}

func TestReadAll(t *testing.T) {
	reader := strings.NewReader(data)
	events := []Event{}
	err := ReadAll(reader, &events)

	if err != nil {
		t.Error(err)
	}

	if len(events) != 2 {
		t.Error("Unexpected results")
	}

	var e Event

	if e = events[0]; e.Id != "0" || e.Name != "Linux" || e.Birthday != "1991" {
		t.Error("Unexpected results")
	}

	if e = events[1]; e.Id != "1" || e.Name != "Apple" || e.Birthday != "1995" {
		t.Error("Unexpected results")
	}
}

func TestReadAllNotPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic")
		}
	}()

	reader := strings.NewReader(data)
	err := ReadAll(reader, false)

	if err != nil {
		t.Error(err)
	}
}

func TestReadAllNotPointerToSlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic")
		}
	}()

	reader := strings.NewReader(data)
	output := "not a slice"
	err := ReadAll(reader, &output)

	if err != nil {
		t.Error(err)
	}
}

func TestReadAllAssignerNotImplemented(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic")
		}
	}()

	reader := strings.NewReader(data)
	events := []string{}
	err := ReadAll(reader, &events)

	if err != nil {
		t.Error(err)
	}
}
