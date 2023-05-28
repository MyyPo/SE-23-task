package utils

import "testing"

func TestFormatNumberWithCommas(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "format to 1,332,444",
			args: args{
				number: 1332444,
			},
			want: "1,332,444",
		},
		{
			name: "format to 993,888",
			args: args{
				number: 993888,
			},
			want: "993,888",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatNumberWithCommas(tt.args.number); got != tt.want {
				t.Errorf("FormatNumberWithCommas() = %v, want %v", got, tt.want)
			}
		})
	}
}
