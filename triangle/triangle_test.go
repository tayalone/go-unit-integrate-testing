//go:build unittest

package triangle

import "testing"

func TestTriagle(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				x: 1,
				y: 0,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				x: 0,
				y: 1,
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				x: 5,
				y: 10,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Triagle(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Triagle() = %v, want %v", got, tt.want)
			}
		})
	}
}
