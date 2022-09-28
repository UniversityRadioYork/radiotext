package common

import "testing"

type testCase struct {
	input    string
	expected string
}

func TestToAscii(t *testing.T) {
	tests := []testCase{
		{
			input:    "Zoë",
			expected: "Zoe",
		},
		{
			input:    "Café",
			expected: "Cafe",
		},
		{
			input:    "Cortège",
			expected: "Cortege",
		},
		{
			input:    "Naïve",
			expected: "Naive",
		},
		{
			input:    "Entrepôt",
			expected: "Entrepot",
		},
		{
			input:    "Façade",
			expected: "Facade",
		},
		{
			input:    "Jalapeño",
			expected: "Jalapeno",
		},
	}

	for _, c := range tests {
		result, err := ToAscii(c.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if result != c.expected {
			t.Errorf("%v became %v not %v", c.input, result, c.expected)
		}
	}
}
