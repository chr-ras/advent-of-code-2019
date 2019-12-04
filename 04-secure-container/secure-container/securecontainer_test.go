package main

import "testing"

func TestHasTwoIdenticalAdjacentDigits(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{
			"123456",
			false,
		},
		{
			"113456",
			true,
		},
		{
			"122456",
			true,
		},
		{
			"123356",
			true,
		},
		{
			"123446",
			true,
		},
		{
			"123455",
			true,
		},
		{
			"111456",
			true,
		},
		{
			"111156",
			true,
		},
		{
			"111116",
			true,
		},
		{
			"111111",
			true,
		},
	}

	for _, c := range cases {
		actual := hasAtleastTwoIdenticalAdjacentDigits(c.input)

		if actual != c.expected {
			t.Errorf("hasTwoIdenticalAdjacentDigits(\"%v\") == %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestHasGroupOfExactlyTwoIdenticalDigits(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{
			"123456",
			false,
		},
		{
			"113456",
			true,
		},
		{
			"122456",
			true,
		},
		{
			"123356",
			true,
		},
		{
			"123446",
			true,
		},
		{
			"123455",
			true,
		},
		{
			"111456",
			false,
		},
		{
			"111156",
			false,
		},
		{
			"111116",
			false,
		},
		{
			"111111",
			false,
		},
		{
			"111122",
			true,
		},
		{
			"111233",
			true,
		},
	}

	for _, c := range cases {
		actual := hasGroupOfExactlyTwoIdenticalDigits(c.input)

		if actual != c.expected {
			t.Errorf("hasGroupOfExactlyTwoIdenticalDigits(\"%v\") == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
