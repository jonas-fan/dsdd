package diagnostic

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	PID         string
	PPID        string
	User        string
	UserID      string
	Path        string
	CommandLine string
	Name        string
}

type Tasks []Task

func (t *Task) assign(name string, value string) {
	switch strings.ToLower(name) {
	case "identifier":
		t.PID = value
	case "parent":
		t.PPID = value
	case "user":
		t.User = value
	case "userid":
		t.UserID = value
	case "path":
		t.Path = value
	case "commandline":
		t.CommandLine = value
	case "process":
		t.Name = value
	}
}

func (t *Task) done() {
	if t.CommandLine == "" {
		t.CommandLine = fmt.Sprintf("[%s]", t.Name)
	}

	if t.Path == "" {
		t.Path = t.CommandLine
	}
}

func (s *Tasks) extend() {
	*s = append(*s, Task{})
}

func (s *Tasks) shrink() {
	*s = (*s)[:len(*s)-1]
}

func (s *Tasks) last() *Task {
	if len(*s) == 0 {
		return nil
	}

	return &(*s)[len(*s)-1]
}

func (s *Tasks) UnmarshalXMLStartElement(start xml.StartElement) {
	task := s.last()

	switch start.Name.Local {
	case "HostMetaData":
		for _, attr := range start.Attr {
			task.assign(attr.Name.Local, attr.Value)
		}
	case "Attribute":
		var name string
		var value string

		for _, attr := range start.Attr {
			switch strings.ToLower(attr.Name.Local) {
			case "name":
				name = attr.Value
			case "value":
				value = attr.Value
			}
		}

		task.assign(name, value)
	}
}

func (s *Tasks) UnmarshalXMLEndElement(end xml.EndElement) {
	task := s.last()

	switch end.Name.Local {
	case "HostMetaData":
		task.done()
		s.extend()
	}
}

func (s *Tasks) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	s.extend()

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.StartElement:
			s.UnmarshalXMLStartElement(element)
		case xml.EndElement:
			s.UnmarshalXMLEndElement(element)
		}
	}

	s.shrink()

	return nil
}

// ReadProcess returns the running processes from a reader.
func ReadProcess(reader io.Reader) ([]Task, error) {
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	var tasks Tasks = make([]Task, 0, 1024)

	if err = xml.Unmarshal(content, &tasks); err != nil {
		return nil, err
	}

	sort.Slice(tasks, func(lhs int, rhs int) bool {
		lhsPID, _ := strconv.Atoi(tasks[lhs].PID)
		rhsPID, _ := strconv.Atoi(tasks[rhs].PID)

		return lhsPID < rhsPID
	})

	return tasks, nil
}
