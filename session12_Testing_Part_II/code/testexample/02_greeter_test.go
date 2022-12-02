package testspart1

import (
	"testing"
)

func TestGreetingTableDriven(t *testing.T) {
	type args struct {
		name string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"empty_input": {
			args: args{
				name: "",
			},
			want: "",
		},
		"non-empty_input": {
			args: args{
				name: "Sergio",
			},
			want: "Greeting, Sergio !",
		},
	}
	for name, tt := range tests {
		tt := tt // PLEASE DO NOT FORGET TO ADD THIS!
		t.Run(name, func(t *testing.T) {
			if got := Greeting(tt.args.name); got != tt.want {
				t.Errorf("Greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO Try to generate table-driven tests for Greeting() using your IDE.
