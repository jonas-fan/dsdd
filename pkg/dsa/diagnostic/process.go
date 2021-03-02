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

type HostMetaData struct {
	PID         string
	PPID        string
	User        string
	UserID      string
	Path        string
	Process     string
	CommandLine string
}

func (m *HostMetaData) Assign(name string, value string) {
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
	case "process":
		m.Process = value
	case "commandline":
		m.CommandLine = value
	}
}

func (m *HostMetaData) Tidy() {
	if m.CommandLine == "" {
		m.CommandLine = fmt.Sprintf("[%s]", m.Process)
	}

	if m.Path == "" {
		m.Path = m.CommandLine
	}
}

type HostMetaDatas struct {
	Data []HostMetaData
}

func (m *HostMetaDatas) UnmarshalXMLStartElement(start xml.StartElement, data *HostMetaData) {
	switch start.Name.Local {
	case "HostMetaData":
		for _, attr := range start.Attr {
			data.Assign(attr.Name.Local, attr.Value)
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

		data.Assign(name, value)
	}
}

func (m *HostMetaDatas) UnmarshalXMLEndElement(end xml.EndElement, data *HostMetaData) {
	switch end.Name.Local {
	case "HostMetaData":
		data.Tidy()
		m.Data = append(m.Data, *data)
		*data = HostMetaData{}
	}
}

func (m *HostMetaDatas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data HostMetaData

	for {
		token, err := d.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.StartElement:
			m.UnmarshalXMLStartElement(element, &data)
		case xml.EndElement:
			m.UnmarshalXMLEndElement(element, &data)
		}
	}

	return nil
}

func readRunningProcess(r io.Reader) ([]HostMetaData, error) {
	content, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	metadatas := HostMetaDatas{
		Data: make([]HostMetaData, 0, 1024),
	}

	if err = xml.Unmarshal(content, &metadatas); err != nil {
		return nil, err
	}

	return metadatas.Data, nil
}

// ReadProcess returns the running processes mentioned in a diagnostic package.
func ReadProcess() ([]HostMetaData, error) {
	name := filepath.Join("Agent/RunningProcesses.xml")
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return readRunningProcess(file)
}
