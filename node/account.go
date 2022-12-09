package node

// Account defines the implementation of an Evmos account
// in the context of the testnode tool.
type Account struct {
	// bech32Address is the account address in bech32 format
	bech32Address string

	// hexAddress is the account address in hex format
	hexAddress string

	// balance is the account's balance
	balance string
}
