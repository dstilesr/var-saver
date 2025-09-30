package confighandle

// PrintableItem are items that can be printed to the CLI.
type PrintableItem interface {
	FormatString() string
}

// Metadata represents metadata of the application and config file.
type Metadata struct {
	Version     string `toml:"app-version"`
	LastUpdated string `toml:"file-last-updated"`
	AppName     string `toml:"app-name"`
}

// Variable represents a stored variable with a name and value
type Variable struct {
	Name        string `toml:"name"`
	Value       string `toml:"value"`
	Environment string `toml:"env"`
}

// Project contains all the variables associated with a given project
type Project struct {
	Name      string      `toml:"name"`
	Variables []*Variable `toml:"variable"`
}

// SavedVariables represents all saved contents in the toml config file.
type SavedVariables struct {
	Meta     *Metadata           `toml:"meta"`
	Projects map[string]*Project `toml:"project"`
}
