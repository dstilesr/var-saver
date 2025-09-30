package confighandle

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml/v2"
)

// GetMetadata Instantiates a metadata object
func GetMetadata() *Metadata {
	curr_time := time.Now().UTC().Format(time.RFC3339)
	return &Metadata{
		LastUpdated: curr_time,
		AppName:     AppName,
		Version:     Version,
	}
}

// NewConfig Instantiates a new configuration from scratch
func NewConfig() *SavedVariables {
	cfg := SavedVariables{
		Projects: make(map[string]*Project),
		Meta:     GetMetadata(),
	}
	return &cfg
}

// FromToml parses the configuration from a toml string
func FromToml(raw []byte) (*SavedVariables, error) {
	sv := &SavedVariables{}
	decoder := toml.NewDecoder(bytes.NewReader(raw))
	err := decoder.Decode(sv)
	if err != nil {
		return nil, err
	}
	return sv, nil
}

// ConfigPath returns the path to the config file. It will be stored at
// `{HOME}/.var-saver/vars.toml`
func ConfigPath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Unable to get user's home directory", "error", err)
		return "", err
	}
	fp := filepath.Join(
		home,
		".var-saver",
		"vars.toml",
	)

	return fp, nil
}
