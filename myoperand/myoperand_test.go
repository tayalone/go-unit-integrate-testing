package myoperand

import "testing"

func TestMyOperand(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
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
			if got := MyOperand(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("MyOperand() = %v, want %v", got, tt.want)
			}
		})
	}
}
