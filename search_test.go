package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		q string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "search",
			args: args{
				q: "gomemcache",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search(tt.args.q)
		})
	}
}
