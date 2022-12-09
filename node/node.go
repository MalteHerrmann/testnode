// Package node contains the Evmos node implementation
// which is used in the testnode tool.
package node

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
