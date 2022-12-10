// Package fs contains utilities for storing and accessing
// the file and folder structure of an Evmos node.
package fs

import "path"

// EvmosFS is the type to describe the file and folder structure
// of an Evmos node.
type EvmosFS struct {
	// homeDir defines where the configuration and data
	// for an Evmos node is stored
	homeDir string

	// genesisFile is the path of the genesis.json file.
	genesisFile string
}

// NewEvmosFS creates a new EvmosFS instance and populates it with
// the given attributes.
func NewEvmosFS(homeDir string) *EvmosFS {
	genesisFile := path.Join(homeDir, "config", "genesis.json")

	return &EvmosFS{
		homeDir:     homeDir,
		genesisFile: genesisFile,
	}
}
