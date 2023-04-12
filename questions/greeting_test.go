package questions

import (
	"testing"
)

func Test_isCatImageTime(t *testing.T) {
	type args struct {
		hourInput      int
		expectedResult bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test 13 hour",
			args: args{
				hourInput:      13,
				expectedResult: false,
			}},
		{name: "test 14 hour",
			args: args{
				hourInput:      14,
				expectedResult: true,
			}},
		{name: "test 15 hour",
			args: args{
				hourInput:      15,
				expectedResult: true,
			}},
		{name: "test 16 hour",
			args: args{
				hourInput:      16,
				expectedResult: false,
			}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := isCatImageTime(test.args.hourInput); result != test.args.expectedResult {
				t.Errorf("isCatImageTime %s test failed!", test.name)
			}
		})

	}
}

func Test_getGreetingText(t *testing.T) {
	type args struct {
		nameInput      string
		expectedResult string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "no name",
			args: args{
				nameInput:      "",
				expectedResult: "hello world",
			}},
		{name: "valid name",
			args: args{
				nameInput:      "Israel",
				expectedResult: "hello Israel",
			}},
		{name: "blanc name",
			args: args{
				nameInput:      "  ",
				expectedResult: "hello world",
			}},
		{name: "spaced name",
			args: args{
				nameInput:      "  Av i ",
				expectedResult: "hello Av i",
			}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := getGreetingText(test.args.nameInput); string(result) != test.args.expectedResult {
				t.Errorf("getGreetingText %s test failed!, expectedRes: %s, actualRes: %s", test.name, test.args.expectedResult, result)
			}
		})

	}
}
