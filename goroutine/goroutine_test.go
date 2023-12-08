package goroutine

import "testing"

func Test_printOddAndEven(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				start: 1,
				end:   100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printOddAndEven(tt.args.start, tt.args.end)
		})
	}
}
