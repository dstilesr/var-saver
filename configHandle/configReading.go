package confighandle

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// VariableExists Determines if a variable already exists in the config
func (sv *SavedVariables) VariableExists(proj, name, env string) bool {
	proj = strings.ToLower(proj)
	projElem, ok := sv.Projects[proj]

	if !ok {
		return false
	}

	for _, v := range projElem.Variables {
		if v.Environment == strings.ToLower(env) && v.Name == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// ReadConfig reads the variables from the saved config file.
// If the file does not exist, creates the parent dir and returns a new
// config.
func ReadConfig() *SavedVariables {
	fp, err := ConfigPath()
	if err != nil {
		panic("Unable to obtain config file path!")
	}

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		parent := filepath.Dir(fp)
		_, err = os.Stat(parent)
		if os.IsNotExist(err) {
			// No public permissions for dir
			os.Mkdir(parent, 0760)
		}
		return NewConfig()
	}

	fc, err := os.ReadFile(fp)
	if err != nil {
		slog.Error("Unable to read config file", "error", err)
		panic("could not read cfg file")
	}

	data, err := FromToml(fc)
	if err != nil {
		slog.Error("Unable to parse config", "error", err)
		panic("could not parse cfg file")
	}
	return data
}

// GetVariable gets a variable from the project. Returns error if the Variable
// does not exist.
func (p *Project) GetVariable(name, env string) (*Variable, error) {
	for _, v := range p.Variables {
		if v.Environment == strings.ToLower(env) && v.Name == strings.ToLower(name) {
			return v, nil
		}
	}
	return nil, fmt.Errorf("variable '%s' not found for env '%s'", name, env)
}

// GetVariable gets a variable from the config. Returns error if the Variable
// does not exist, or if the project does not exist.
func (sv *SavedVariables) GetVariable(project, name, env string) (*Variable, error) {
	project = strings.ToLower(project)
	p, ok := sv.Projects[project]
	if !ok {
		return nil, fmt.Errorf("project '%s' not found", project)
	}
	return p.GetVariable(name, env)
}
