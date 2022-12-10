// Package node contains the Evmos node implementation
// which is used in the testnode tool.
package node

var DefaultNames = []string{"dev0", "dev1", "dev2"}

// EvmosNode is the implementation of an Evmos node for the
// testnode tool.
type EvmosNode struct {
	// url is the host url of the Evmos node
	url string

	// evmPort is the port id at which the EVM requests are received
	evmPort int

	// accounts is a slice of Account types
	accounts []Account
}

// NewEvmosNode creates a new EvmosNode instance and populates it with
// the given attributes.
func NewEvmosNode(url string, evmPort int) *EvmosNode {
	var accounts []Account //nolint:prealloc

	for _, name := range DefaultNames {
		newAcc := NewAccount(name)
		accounts = append(accounts, newAcc)
	}

	return &EvmosNode{
		url:      url,
		evmPort:  evmPort,
		accounts: accounts,
	}
}
