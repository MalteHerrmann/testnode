package node

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEvmosNode(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		evmPort int
		expKeys []string
	}{
		{
			name:    "default Evmos node",
			url:     "http://localhost",
			evmPort: 8545,
			expKeys: DEFAULT_NAMES,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := NewEvmosNode(tt.url, tt.evmPort)
			require.Equal(t, returned.url, tt.url)
			require.Equal(t, returned.evmPort, tt.evmPort)

			for idx, acc := range returned.accounts {
				require.Equal(t, acc.name, tt.expKeys[idx])
			}
		})
	}
}
