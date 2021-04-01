package diagnostic

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

func (m *Task) Assign(name string, value string) {
	switch strings.ToLower(name) {
	case "identifier":
		m.PID = value
	case "parent":
		m.PPID = value
	case "user":
		m.User = value
	case "userid":
		m.UserID = value
	case "path":
		m.Path = value
	case "commandline":
		m.CommandLine = value
	case "process":
		m.Name = value
	}
}

func (m *Task) Tidy() {
	if m.CommandLine == "" {
		m.CommandLine = fmt.Sprintf("[%s]", m.Name)
	}

	if m.Path == "" {
		m.Path = m.CommandLine
	}
}

func (t *Tasks) UnmarshalXMLStartElement(task *Task, start xml.StartElement) {
	switch start.Name.Local {
	case "HostMetaData":
		for _, attr := range start.Attr {
			task.Assign(attr.Name.Local, attr.Value)
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

		task.Assign(name, value)
	}
}

func (t *Tasks) UnmarshalXMLEndElement(task *Task, end xml.EndElement) {
	switch end.Name.Local {
	case "HostMetaData":
		task.Tidy()
		*t = append(*t, *task)
		*task = Task{}
	}
}

func (t *Tasks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	task := Task{}

	for {
		token, err := d.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.StartElement:
			t.UnmarshalXMLStartElement(&task, element)
		case xml.EndElement:
			t.UnmarshalXMLEndElement(&task, element)
		}
	}

	return nil
}

func readProcess(r io.Reader) ([]Task, error) {
	content, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var tasks Tasks = make([]Task, 0, 1024)

	if err = xml.Unmarshal(content, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// ReadProcess returns the running processes mentioned from a location.
func ReadProcessFrom(filename string) ([]Task, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return readProcess(file)
}

// ReadProcess returns the running processes mentioned in a diagnostic package.
func ReadProcess() ([]Task, error) {
	return ReadProcessFrom(filepath.Join("Agent", "RunningProcesses.xml"))
}
