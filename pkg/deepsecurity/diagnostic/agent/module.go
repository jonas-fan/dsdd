package agent

import (
	"encoding/json"
	"regexp"
	"sort"
	"strings"
)

type Module struct {
	Name       string
	Type       string
	Version    string
	Dependency []*Module
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

func makeModule(name string, info map[string]interface{}) *Module {
	module := &Module{
		Name: name,
	}

	if value, ok := info["type"]; ok {
		module.Type = value.(string)
	} else {
		module.Type = "plugin"
	}

	if value, ok := info["version"]; ok {
		module.Version = strings.ReplaceAll(value.(string), "-", ".")
	}

	if value, ok := info["deps"]; ok {
		module.Dependency = makeModules(value.(map[string]interface{}))
	}

	return module
}

func makeModules(jdata map[string]interface{}) []*Module {
	modules := make([]*Module, 0, len(jdata))

	for name, info := range jdata {
		module := makeModule(name, info.(map[string]interface{}))
		modules = append(modules, module)
	}

	sort.Slice(modules, func(lhs int, rhs int) bool {
		return modules[lhs].compare(modules[rhs])
	})

	return modules
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
