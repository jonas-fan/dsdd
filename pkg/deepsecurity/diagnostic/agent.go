package diagnostic

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Agent struct {
	Snapshot string
	Os       System
	Platform string
	Version  string
	Module   []Module
	Guid     Guid
}

type System struct {
	Type   string
	Arch   string
	Kernel string
}

type Module struct {
	Name       string
	Type       string
	Version    string
	Dependency []Module
}

type Guid struct {
	XMLName xml.Name `xml:"guids"`
	Agent   string   `xml:"AgentGUID"`
	Manager string   `xml:"DSMGUID"`
}

func (a *Agent) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "dsa.snapshot = %s", a.Snapshot)
	fmt.Fprintf(&builder, "\ndsa.os = %s", a.Os.Type)
	fmt.Fprintf(&builder, "\ndsa.os.arch = %s", a.Os.Arch)
	fmt.Fprintf(&builder, "\ndsa.os.kernel = %s", a.Os.Kernel)
	fmt.Fprintf(&builder, "\ndsa.platform = %s", a.Platform)
	fmt.Fprintf(&builder, "\ndsa.version = %s", a.Version)
	fmt.Fprintf(&builder, "\ndsa.guid = %s", a.Guid.Agent)
	fmt.Fprintf(&builder, "\ndsa.manager.guid = %s", a.Guid.Manager)

	for i, module := range a.Module {
		prefix := fmt.Sprintf("\ndsa.module%d.%s.%s", i, module.Type, module.Name)

		fmt.Fprintf(&builder, "%s = %s", prefix, module.Version)

		for i, dep := range module.Dependency {
			fmt.Fprintf(&builder, "%s.dependency%d = %s", prefix, i, dep.Name)
		}
	}

	return builder.String()
}

func makeModules(jdata map[string]interface{}) []Module {
	var i int
	var modules = make([]Module, len(jdata))

	for name, info := range jdata {
		modules[i].load(name, info.(map[string]interface{}))
		i++
	}

	sort.Slice(modules, func(lhs int, rhs int) bool {
		return modules[lhs].compare(&modules[rhs])
	})

	return modules
}

func (m *Module) load(name string, info map[string]interface{}) {
	m.Name = name

	if value, ok := info["type"]; ok {
		m.Type = value.(string)
	} else {
		m.Type = "plugin"
	}

	if value, ok := info["version"]; ok {
		m.Version = strings.ReplaceAll(value.(string), "-", ".")
	}

	if value, ok := info["deps"]; ok {
		m.Dependency = makeModules(value.(map[string]interface{}))
	}
}

func (m *Module) isCore() bool {
	return m.Name == "core" || m.Type == "agent"
}

func (m *Module) isFeature() bool {
	return m.Type == "feature"
}

func (m *Module) isPlugin() bool {
	return m.Type == "plugin"
}

func (m *Module) isDriver() bool {
	return strings.HasSuffix(m.Name, ".ko")
}

func (m *Module) compare(rhs *Module) bool {
	switch {
	case m.isCore() && !rhs.isCore():
		return true
	case !m.isCore() && rhs.isCore():
		return false
	case m.isFeature() && !rhs.isFeature():
		return true
	case !m.isFeature() && rhs.isFeature():
		return false
	case m.isPlugin() && !rhs.isPlugin():
		return true
	case !m.isPlugin() && rhs.isPlugin():
		return false
	case m.Type != rhs.Type:
		return m.Type < rhs.Type
	case !m.isDriver() && rhs.isDriver():
		return true
	case m.isDriver() && !rhs.isDriver():
		return false
	default:
		return m.Name < rhs.Name
	}
}

func toOS(platform string) string {
	platform = strings.ToLower(platform)

	switch {
	case strings.HasPrefix(platform, "aix"):
		return "aix"
	case strings.HasPrefix(platform, "solaris"):
		return "solaris"
	case strings.HasPrefix(platform, "windows"):
		return "windows"
	default:
		return "linux"
	}
}

// This is a trick to convert Lua table to json-like format :)
//
// A Lua table looks like:
//
//     am:{
//         deps = {
//             update = {
//                 type = "Plugin"
//             }
//         },
//         version = "12.0.0-791",
//         type = "feature"
//     }
//     dsa_filter.ko:{
//         version = "12.0.0-791"
//     }
//
func luaToJSON(luatable string) string {
	type rule struct {
		pattern     string
		replacement string
	}

	var data = "{" + strings.ToLower(luatable) + "}"
	var rules = []rule{
		{pattern: `\r\n`, replacement: ``},
		{pattern: `\n`, replacement: ``},
		{pattern: `=`, replacement: `:`},
		{pattern: `\s*([\w.-]+)\s*:`, replacement: `"${1}":`},
		{pattern: `}\s*"`, replacement: `},"`},
	}

	for _, rule := range rules {
		express := regexp.MustCompile(rule.pattern)

		data = express.ReplaceAllString(data, rule.replacement)
	}

	return data
}

func loadJSON(data string) (map[string]interface{}, error) {
	var out map[string]interface{}

	if err := json.Unmarshal([]byte(data), &out); err != nil {
		return nil, err
	}

	return out, nil
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

	var agent = &Agent{}
	var filename string

	filename = filepath.Join(dirname, "timestamp.txt")

	if buf, err := ioutil.ReadFile(filename); err == nil {
		agent.Snapshot = string(buf)
	}

	filename = filepath.Join(dirname, "plugins.txt")

	if blob, err := ioutil.ReadFile(filename); err == nil {
		tokens := strings.SplitN(string(blob), "\r\n", 2)

		if len(tokens) > 1 {
			if jdata, err := loadJSON(luaToJSON(tokens[1])); err == nil {
				agent.Module = makeModules(jdata)
			}
		}

		if tokens = strings.Split(tokens[0], " "); len(tokens) > 1 {
			agent.Version = tokens[1]
		}

		if tokens = strings.Split(tokens[0], ";"); len(tokens) > 5 {
			agent.Platform = tokens[0] + "." + tokens[5]
			agent.Os.Kernel = strings.Join(tokens[1:5], ".")
			agent.Os.Arch = tokens[5]
			agent.Os.Type = toOS(agent.Platform)
		}
	}

	filename = filepath.Join(dirname, "guids.xml")

	if buf, err := ioutil.ReadFile(filename); err == nil {
		xml.Unmarshal(buf, &agent.Guid)
	}

	return agent, nil
}
