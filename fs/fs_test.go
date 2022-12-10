package fs

import (
	"os"
	"reflect"
	"testing"
)

var HOME_DIR = os.Getenv("HOME")

func TestNewEvmosFS(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want *EvmosFS
	}{
		{
			name: "default home directory",
			args: args{
				homeDir: HOME_DIR,
			},
			want: &EvmosFS{
				homeDir:     HOME_DIR,
				genesisFile: HOME_DIR + "/config/genesis.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvmosFS(tt.args.homeDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvmosFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
