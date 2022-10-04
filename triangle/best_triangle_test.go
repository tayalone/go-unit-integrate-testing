package triangle

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *BestTrg
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestTrg_Operand(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		bs   *BestTrg
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
			bs:   New(),
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				x: 0,
				y: 1,
			},
			bs:   New(),
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				x: 5,
				y: 10,
			},
			bs:   New(),
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BestTrg{}
			if got := bs.Operand(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("BestTrg.Operand() = %v, want %v", got, tt.want)
			}
		})
	}
}
