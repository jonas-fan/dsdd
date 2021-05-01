package agent

import (
	"bufio"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity"
)

type Agent struct {
	Version   string
	System    System
	Module    []*Module
	Guid      Guid
	Timestamp string
}

type System struct {
	Name    string
	Type    string
	Release string
}

type Guid struct {
	XMLName xml.Name `xml:"guids"`
	Agent   string   `xml:"AgentGUID"`
	Manager string   `xml:"DSMGUID"`
}

func parsePlatformString(token string) (*System, error) {
	tokens := strings.Split(token, ";")

	if len(tokens) < 6 {
		return nil, errors.New("invalid argument")
	}

	system := &System{
		Name:    tokens[0] + "." + tokens[5],
		Type:    deepsecurity.ToOS(tokens[0]),
		Release: strings.Join(tokens[1:5], "."),
	}

	return system, nil
}

func (a *Agent) parsePlatform(scanner *bufio.Scanner) error {
	if scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		system, err := parsePlatformString(tokens[0])

		if err != nil {
			return err
		}

		a.System = *system
		a.Version = tokens[1]
	} else if err := scanner.Err(); err != nil {
		return err
	} else {
		return errors.New("nothing to read")
	}

	return nil
}

func (a *Agent) parseModule(scanner *bufio.Scanner) error {
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	json := luaToJSON(strings.Join(lines, "\n"))
	jdata, err := loadJSON(json)

	if err != nil {
		return err
	}

	a.Module = makeModules(jdata)

	return nil
}

// ParseInfo retrieves the agent information.
func (a *Agent) ParseInfo(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	a.parsePlatform(scanner)
	a.parseModule(scanner)

	return nil
}

// ParseGuid retrieves the agent and maanger guids.
func (a *Agent) ParseGuid(filename string) error {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = xml.Unmarshal(buf, &a.Guid)

	if err != nil {
		return err
	}

	a.Guid.Agent = strings.ToLower(a.Guid.Agent)
	a.Guid.Manager = strings.ToLower(a.Guid.Manager)

	return nil
}

// ParseTimestamp retrieves the timestamp when collecting diagnostic logs.
func (a *Agent) ParseTimestamp(filename string) error {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	a.Timestamp = string(buf)

	return nil
}

// NewAgent returns a new Agent.
func NewAgent(dirname string) (*Agent, error) {
	dirname, err := filepath.Abs(dirname)

	if err != nil {
		return nil, err
	} else if info, err := os.Stat(dirname); err != nil {
		return nil, err
	} else if !info.IsDir() {
		return nil, errors.New(dirname + ": not a directory")
	}

	agent := &Agent{}
	agent.ParseInfo(filepath.Join(dirname, "plugins.txt"))
	agent.ParseGuid(filepath.Join(dirname, "guids.xml"))
	agent.ParseTimestamp(filepath.Join(dirname, "timestamp.txt"))

	return agent, nil
}
