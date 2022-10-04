package star

import (
	"testing"
)

func TestStar(t *testing.T) {
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
				x: 10,
				y: 0,
			},
			want: 20,
		},
		{
			name: "case 2",
			args: args{
				x: 0,
				y: 10,
			},
			want: 30,
		},
		{
			name: "case 3",
			args: args{
				x: 30,
				y: 10,
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Star(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Star() = %v, want %v", got, tt.want)
			}
		})
	}
}
