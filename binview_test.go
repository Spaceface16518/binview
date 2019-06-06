package binview

import "testing"

func Test_formatBytes(t *testing.T) {
	type args struct {
		input   []byte
		perLine int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				input:   []byte{0, 1, 2, 3, 4, 5, 6, 7},
				perLine: 1,
			},
			want: `00000000
00000001
00000010
00000011
00000100
00000101
00000110
00000111
`,
		},
		{
			name: "4",
			args: args{
				input:   []byte{0, 1, 2, 3, 4, 5, 6, 7},
				perLine: 4,
			},
			want: `00000000 00000001 00000010 00000011
00000100 00000101 00000110 00000111
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatBytes(tt.args.input, tt.args.perLine)
			if got != tt.want {
				t.Errorf("formatBytes() = \n%s\nwant \n%s", got, tt.want)
			}
			if len(got) != 9*len(tt.args.input) {
				// TODO: refactor these variables
				t.Errorf("Length did not match predictive algorithm: got %d, want %d", len(got), len(tt.args.input))
			}
		})
	}
}
