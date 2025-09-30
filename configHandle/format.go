package confighandle

import (
	"fmt"
	"strings"
)

// PrintItem Prints an item that can be formatted to the CLI
func PrintItem(i PrintableItem) {
	fmt.Print(i.FormatString())
}

func (m *Metadata) FormatString() string {
	return fmt.Sprintf(metaTemplate, m.AppName, m.Version, m.LastUpdated)
}

func (v *Variable) FormatString() string {
	return fmt.Sprintf(varTemplate, v.Name, v.Environment)
}

func (p *Project) FormatString() string {
	var_strings := make([]string, 0, len(p.Variables))
	for _, v := range p.Variables {
		var_strings = append(var_strings, v.FormatString())
	}
	var_detail := strings.Join(var_strings, "\n")
	return fmt.Sprintf(projectTemplate, p.Name, len(p.Variables), var_detail)
}
