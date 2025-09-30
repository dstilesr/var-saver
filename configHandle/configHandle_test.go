package confighandle

import "testing"

// TestGetProject tests the functionality to get / create a project in the config
func TestGetProject(t *testing.T) {
	sv := NewConfig()
	var p1, p2, p3, p4 string = "proj1", "PROj1", "proj3", "pRoJ4"

	sv.GetProject(p1)
	if len(sv.Projects) != 1 {
		t.Error("Did not create a single project correctly")
	}

	sv.GetProject(p2)
	if len(sv.Projects) != 1 {
		t.Error("Did not match project name when letter case changed")
	}
	_, ok := sv.Projects[p1]
	if !ok {
		t.Error("Project name not stored in lowercase")
	}

	sv.GetProject(p3)
	po := sv.GetProject(p4)

	if len(sv.Projects) != 3 {
		t.Error("Did not create expected number of projects")
	}

	if len(po.Variables) != 0 {
		t.Error("New project has non-empty var list")
	}
}

// TestGetVar Tests project functionality to get a variable
func TestGetVar(t *testing.T) {
	p := &Project{
		Variables: []*Variable{
			&Variable{Name: "v1", Value: "val1", Environment: "env1"},
			&Variable{Name: "v1", Value: "val2", Environment: "env2"},
			&Variable{Name: "v2", Value: "val3", Environment: "env1"},
		},
	}

	v, err := p.GetVariable("v1", "env1")
	if v == nil || err != nil {
		t.Error("Failed to get existing variable")
	}
	if v.Environment != "env1" || v.Name != "v1" || v.Value != "val1" {
		t.Error("Incorrect variable properties")
	}

	v, err = p.GetVariable("v1", "env2")
	if v == nil || err != nil {
		t.Error("Failed to get existing variable from other env")
	}
	if v.Environment != "env2" || v.Name != "v1" || v.Value != "val2" {
		t.Error("Incorrect variable properties - other env")
	}

	v, err = p.GetVariable("v6", "env1")
	if err == nil {
		t.Error("Retrieved nonexistent variable")
	}

	v, err = p.GetVariable("V1", "ENV2")
	if v == nil || err != nil {
		t.Error("Failed to get existing variable from other env - cased ")
	}
	if v.Environment != "env2" || v.Name != "v1" || v.Value != "val2" {
		t.Error("Incorrect variable properties - other env - cased")
	}
}

// TestAddVar tests the functionality to add a new variable
func TestAddVar(t *testing.T) {
	sv := NewConfig()

	err := sv.AddVariable("proj1", "var1", "env", "Value-1", false)
	if err != nil {
		t.Error("Could not add variable to empty config")
	}
	if len(sv.Projects) != 1 {
		t.Error("Did not add new project for variable")
	}

	p, ok := sv.Projects["proj1"]
	if !ok {
		t.Error("Project does not have expected name")
	}
	if len(p.Variables) != 1 {
		t.Error("Did not add new variable to project")
	}

	v := p.Variables[0]
	if v.Name != "var1" || v.Environment != "env" || v.Value != "Value-1" {
		t.Error("Variable not created with correct properties")
	}

	// Overwrite - failed
	err = sv.AddVariable("proj1", "var1", "env", "Value-2", false)
	if err == nil {
		t.Error("Did not raise error when overwriting variable")
	}
	if len(p.Variables) != 1 {
		t.Error("Incorrect variable count after failed overwrite")
	}
	v = p.Variables[0]
	if v.Name != "var1" || v.Environment != "env" || v.Value != "Value-1" {
		t.Error("Variable updated with overwrite set to false")
	}

	// Overwrite - success
	err = sv.AddVariable("proj1", "var1", "env", "Value-2", true)
	if err != nil {
		t.Error("Failed to overwrite variable")
	}
	if len(p.Variables) != 1 {
		t.Error("Incorrect variable count after overwrite")
	}
	v = p.Variables[0]
	if v.Name != "var1" || v.Environment != "env" || v.Value != "Value-2" {
		t.Error("Incorrect variable properties after overwrite")
	}

	// Overwrite - success - cases differ
	err = sv.AddVariable("proj1", "VAR1", "ENV", "Value-3", true)
	if err != nil {
		t.Error("Failed to overwrite variable")
	}
	if len(p.Variables) != 1 {
		t.Error("Incorrect variable count after overwrite")
	}
	v = p.Variables[0]
	if v.Name != "var1" || v.Environment != "env" || v.Value != "Value-3" {
		t.Error("Incorrect variable properties after overwrite - cased")
	}

	// Same name, different env
	err = sv.AddVariable("proj1", "var1", "env2", "Value-4", false)
	if err != nil {
		t.Error("Failed to add variable to new env")
	}
	if len(p.Variables) != 2 {
		t.Error("Incorrect variable count after adding new")
	}
	v, err = p.GetVariable("var1", "env2")
	if err != nil {
		t.Error("Could not retrieve new variable")
	}
	if v.Name != "var1" || v.Environment != "env2" || v.Value != "Value-4" {
		t.Error("Incorrect variable properties")
	}

	// Old variable should not have changed
	v, err = p.GetVariable("var1", "env")
	if err != nil {
		t.Error("Could not retrieve old variable")
	}
	if v.Name != "var1" || v.Environment != "env" || v.Value != "Value-3" {
		t.Error("Incorrect variable properties after adding new")
	}
}
