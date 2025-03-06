package util

import (
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

func TestStringFlag(t *testing.T) {
	type args struct {
		env   string
		name  string
		usage string
		value string
	}
	tests := []struct {
		name string
		args args
		want cli.StringFlag
	}{
		{
			name: "success case",
			args: args{
				env:   "TEST_ENV",
				name:  "TEST_NAME",
				usage: "TEST_USAGE",
				value: "TEST_VALUE",
			},
			want: cli.StringFlag{
				EnvVar: "TEST_ENV",
				Name:   "TEST_NAME",
				Usage:  "TEST_USAGE",
				Value:  "TEST_VALUE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFlag(tt.args.env, tt.args.name, tt.args.usage, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
