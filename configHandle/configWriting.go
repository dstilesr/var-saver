package confighandle

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
)

// Update Sets the metadata's last updated time to the current UTC timestamp
func (m *Metadata) Update() {
	curr_time := time.Now().UTC().Format(time.RFC3339)
	m.LastUpdated = curr_time
}

// SaveCfg saves the configuration to the config file path
func (sv *SavedVariables) SaveCfg() error {
	fp, err := ConfigPath()
	if err != nil {
		return err
	}

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	sv.Meta.Update()
	err = toml.NewEncoder(f).Encode(sv)
	if err != nil {
		slog.Error("Unable to encode TOML to file!", "error", err)
	}
	return err
}

// GetProject gets a project by name from the config. If it does not already
// exists, a new project is created.
func (sv *SavedVariables) GetProject(projName string) *Project {
	nameNorm := strings.ToLower(projName)
	p, ok := sv.Projects[nameNorm]

	if !ok {
		p = &Project{
			Variables: make([]*Variable, 0),
			Name:      nameNorm,
		}
		sv.Projects[nameNorm] = p
	}
	return p
}

// AddVariable Adds a new variable to the configuration. Returns error if the variable exists
// and overwrite is set to false.
func (sv *SavedVariables) AddVariable(proj, name, env, value string, overwrite bool) error {
	if sv.VariableExists(proj, name, env) {
		if !overwrite {
			return fmt.Errorf(
				"the variable '%s' already exists in project '%s' for environment '%s'",
				name,
				proj,
				env,
			)
		}

		// Update existing
		p := sv.GetProject(proj)
		v, err := p.GetVariable(name, env)
		if v == nil || err != nil {
			slog.Error("Could not get existing variable!")
			panic("Could not get existing variable")
		}
		v.Value = value
		return nil
	}

	p := sv.GetProject(proj)
	newVar := Variable{
		Name:        strings.ToLower(name),
		Environment: strings.ToLower(env),
		Value:       value,
	}

	p.Variables = append(p.Variables, &newVar)
	return nil
}

// RemoveVariable Removes a variable from the given project. If the variable
// was not found, returns an error.
func (p *Project) RemoveVariable(name, env string) error {
	new_vars := make([]*Variable, 0, len(p.Variables))
	name = strings.ToLower(name)
	env = strings.ToLower(env)

	removed := false
	for _, v := range p.Variables {
		if v.Name == name && v.Environment == env {
			removed = true
		} else {
			new_vars = append(new_vars, v)
		}
	}

	p.Variables = new_vars
	if !removed {
		return fmt.Errorf(
			"Variable '%s' for environment '%s' not found in the '%s' project",
			name,
			env,
			p.Name,
		)
	}
	return nil
}

// RemoveProject Deletes a project from the config along with all its
// associated variables. Returns error if the project does not exist.
func (sv *SavedVariables) RemoveProject(name string) error {
	name = strings.ToLower(name)
	_, ok := sv.Projects[name]
	if !ok {
		return fmt.Errorf("Project '%s' not found", name)
	}
	delete(sv.Projects, name)
	return nil
}

// RemoveVariable Removes a variable from a project in the config
func (sv *SavedVariables) RemoveVariable(project, name, env string) error {
	name = strings.ToLower(name)
	project = strings.ToLower(project)
	p, ok := sv.Projects[project]
	if !ok {
		return fmt.Errorf("Project '%s' not found", name)
	}
	return p.RemoveVariable(name, env)
}
