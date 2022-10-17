package analyse

import (
	"testing"
)

func Test_hasDoubleLetters(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "double letter",
			args: args{
				word: "floor",
			},
			want: true,
		},
		{
			name: "no double letter",
			args: args{
				word: "vodka",
			},
			want: false,
		},
		{
			name: "empty string",
			args: args{
				word: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDoubleLetters(tt.args.word); got != tt.want {
				t.Errorf("hasDoubleLetters(%s) = %v, want %v", tt.args.word, got, tt.want)
			}
		})
	}
}

func Fuzz_hasDoubleLetters(f *testing.F) {
	testcases := []string{"floor", "vodka", ""}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, w string) {
		d := hasDoubleLetters(w)
		if !d && d {
			t.Errorf("The word %s was a problem for hasDoubleLetters()", w)
		}
	})
}

func TestNumDoubleLetters(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a few words",
			args: args{
				words: []string{"floor", "vodka", ""},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDoubleLetters(tt.args.words); got != tt.want {
				t.Errorf("NumDoubleLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
