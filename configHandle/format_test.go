package confighandle

import "testing"

// TestRender tests that templates render without errors (would cause panic)
func TestRender(t *testing.T) {
	sv := GetTestConfigObj()

	p := sv.Projects["proj1"]
	p.FormatString()

	p.Variables[0].FormatString()

	sv.Meta.FormatString()
}
