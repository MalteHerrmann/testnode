package node

// Account defines the implementation of an Evmos account
// in the context of the testnode tool.
type Account struct {
	// // balance is the account's balance
	// balance string
	// // bech32Address is the account address in bech32 format
	// bech32Address string
	// // hexAddress is the account address in hex format
	// hexAddress string
	// name is the key name
	name string
}

// NewAccount creates a new Account and populates it with the
// given name.
func NewAccount(name string) Account {
	return Account{
		name: name,
	}
}
