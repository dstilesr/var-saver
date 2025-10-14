package confighandle

import "testing"

// TestFromToml tests parsing config from TOML strings
func TestFromToml(t *testing.T) {
	test_toml := `
[meta]
app-version = "v0.0.1"
file-last-updated = "2020-01-01T12:01:15Z"
app-name = "var-saver"

[project.proj1]
name = "proj1"
description = "This is project proj1"

[[project.proj1.variable]]
name = "v1"
env = "env1"
value = "val1"

[[project.proj1.variable]]
name = "v1"
env = "env2"
value = "val2"

[[project.proj1.variable]]
name = "v2"
env = "env2"
value = "val3"
`
	sv, err := FromToml([]byte(test_toml))
	if err != nil {
		t.Errorf("Failed to parse toml string. Error: %v", err)
	}

	if len(sv.Projects) != 1 {
		t.Error("Parsed wrong number of projects")
	}
	p, ok := sv.Projects["proj1"]
	if !ok {
		t.Error("Did not find expected project")
	}

	if p.Description != "This is project proj1" {
		t.Error("Did not parse expected project description")
	}

	if len(p.Variables) != 3 {
		t.Error("Parsed project with unexpected number of variables")
	}

	if sv.Meta.Version != "v0.0.1" || sv.Meta.AppName != "var-saver" {
		t.Log(sv.Meta)
		t.Error("Incorrect metadata fields parsed")
	}

	if sv.Meta.LastUpdated != "2020-01-01T12:01:15Z" {
		t.Log(sv.Meta)
		t.Error("Incorrect metadata fields parsed")
	}
}
