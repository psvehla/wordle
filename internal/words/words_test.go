package words

import (
	"testing"
)

func TestGetWords(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "test",
			want:  12972,
			want1: 2309,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetWords()
			if len(got) != tt.want {
				t.Errorf("len(GetWords()) got = %v, want %v", got, tt.want)
			}
			if len(got1) != tt.want1 {
				t.Errorf("len(GetWords()) got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
