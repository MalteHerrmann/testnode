package node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name    string
		keyName string
	}{
		{
			name:    "new account",
			keyName: "dev0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := NewAccount(tt.name)
			require.Equal(t, acc.name, tt.name)
		})
	}
}
