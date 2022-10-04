package myoperand

import (
	"reflect"
	"testing"

	"github.com/tayalone/go-unit-integrate-test/star"
	"github.com/tayalone/go-unit-integrate-test/triangle"
)

type MOP struct{}

var mop = MOP{}

func (mop *MOP) Operand(x int, y int) int {
	if x == 30 && y == 10 {
		return 27
	}
	if x == 10 && y == 5 {
		return 25
	}
	return 1
}

func TestNew(t *testing.T) {
	type args struct {
		star     star.StrItf
		triangle triangle.TrgItf
	}
	tests := []struct {
		name string
		args args
		want *MoStr
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				star:     &mop,
				triangle: &mop,
			},
			want: New(&mop, &mop),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.star, tt.args.triangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoStr_Operand(t *testing.T) {
	type fields struct {
		star     star.StrItf
		triangle triangle.TrgItf
	}
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			fields: fields{
				star:     &mop,
				triangle: &mop,
			},
			args: args{
				a: 30,
				b: 10,
				c: 10,
				d: 5,
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mo := &MoStr{
				star:     tt.fields.star,
				triangle: tt.fields.triangle,
			}
			if got := mo.Operand(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("MoStr.Operand() = %v, want %v", got, tt.want)
			}
		})
	}
}
