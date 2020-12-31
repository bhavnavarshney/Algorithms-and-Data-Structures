package strings

import "testing"

func TestKMP(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				haystack: "aabaaabaaac",
				needle:   "aabaaac",
			},
			want: 4,
		},
		{
			name: "test-2",
			args: args{
				haystack: "onionionions",
				needle:   "onions",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
