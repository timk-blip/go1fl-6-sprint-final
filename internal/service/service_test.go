package service

import "testing"

func TestIsParseable(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestIsParseableMorse",
			args: args{".--. .-. .. .-- . -"},
			want: ".--. .-. .. .-- . -",
		},
		{
			name: "TestIsParseableMorseErr",
			args: args{".--. .-. .. .-- .12"},
			want: "",
		},
		{
			name: "TestIsParseableMorseErr1",
			args: args{"df2. .-. .. .-- .12"},
			want: "",
		},
		{
			name: "TestIsParseableMorse1",
			args: args{".--. .-. .. .-- . -"},
			want: ".--. .-. .. .-- . -",
		},
		{
			name: "TestIsParseableText",
			args: args{"привет Привет"},
			want: "привет Привет",
		},
		{
			name: "TestIsParseableTextErr",
			args: args{"Привет .--. .-. .. .-- . -"},
			want: "",
		},
		{
			name: "TestIsParseableErr",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParseable(tt.args.text); got != tt.want {
				t.Errorf("IsParseable() = %v, want %v", got, tt.want)
			}
		})
	}
}
