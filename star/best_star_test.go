package star

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *BestStr
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			want: (*BestStr)(New()),
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

func TestBestStr_Operand(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		bs   *BestStr
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			bs:   New(),
			args: args{
				x: 10,
				y: 0,
			},
			want: 20,
		},
		{
			name: "case 2",
			bs:   New(),
			args: args{
				x: 0,
				y: 10,
			},
			want: 30,
		},
		{
			name: "case 3",
			bs:   New(),
			args: args{
				x: 30,
				y: 10,
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BestStr{}
			if got := bs.Operand(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("BestStr.Operand() = %v, want %v", got, tt.want)
			}
		})
	}
}
