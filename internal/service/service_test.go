package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsParseable1(t *testing.T) {
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
			want: "ПРИВЕТ",
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
			want: "ПРИВЕТ",
		},
		{
			name: "TestIsParseableText",
			args: args{"НХТТГНЩЧКОФВЙЭМПХБСЭПАЩ"},
			want: "-. .... - - --. -. --.- ---. -.- --- ..-. .-- .--- ..-.. -- .--. .... -... ... ..-.. .--. .- --.-",
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
			got := IsParseable(tt.args.text)
			assert.Equal(t, tt.want, got, "Ошибка в тест-кейсе: %s", tt.name)
		})
	}
}
