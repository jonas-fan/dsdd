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
		t.Errorf("Expected length: 2, got: %d", len(events))
	}

	var expected = []Event{
		{Id: "0", Name: "Linux", Birthday: "1991"},
		{Id: "1", Name: "Apple", Birthday: "1995"},
	}

	for index := range events {
		if events[index] != expected[index] {
			t.Errorf("Expected: %s, got: %s", expected[index], events[index])
		}
	}
}

func TestReadAllNotPointer(t *testing.T) {
	reader := strings.NewReader(data)
	err := ReadAll(reader, false)

	if err == nil {
		t.Errorf("Expected an error, got: nil")
	}
}

func TestReadAllNotPointerToSlice(t *testing.T) {
	reader := strings.NewReader(data)
	output := "not a slice"
	err := ReadAll(reader, &output)

	if err == nil {
		t.Errorf("Expected an error, got: nil")
	}
}

func TestReadAllAssignerNotImplemented(t *testing.T) {
	reader := strings.NewReader(data)
	events := []string{}
	err := ReadAll(reader, &events)

	if err == nil {
		t.Errorf("Expected an error, got: nil")
	}
}
