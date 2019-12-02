package num2words

import (
	"testing"
	"strings"
)

func TestZero( t *testing.T) {
	if ToWords(0) !="zero"{
		t.Log("0 is not zero")
		t.Fail()
	}
}

func TestThirteen(t *testing.T) {
	if ToWords(13) != "thirteen" {
		t.Log("not thirteen")
		t.Fail()
	}
}

func TestTwoHundredThirtyNine ( t *testing.T) {
	if ToWords(239) != "two hundred thirty-nine" {
		t.Log("Not 239")
		t.Fail()
	}
}

func TestOneThousandOneHundredEleven ( t *testing.T) {
	if ToWords(1111) != "one thousand one hundred eleven" {
		t.Log("Not 1111")
		t.Fail()
	}
} 

func TestNumbers(t *testing.T) {
	/*
	testcase := [...]string {"zero", "one", "two", "three"}

	for index, value := range testcase {
		if r := ToWords(index); r != value {
			t.Logf("Input: %d, Expected: %s, Got: %s\n", index, value, r)
			t.Fail()
		}
	}
	*/
	testcase := map[int]string {
		0: "zero",
		1: "one",
		100: "one hundred",
		1111: "one thousand one hundred eleven",
		6738: "six thousand seven hundred thirty-eight",
	}

	for index, value := range testcase {
		if r := strings.Trim(ToWords(index), " "); r != value {
			t.Logf("Input: %d, Expected: %s, Got: %s\n", index, value, r)
			t.Fail()
		}
	}
}

/* 
t.Log() - to mark about test
t.Logf() - supports formatted log
t.Fail() - If test fails, then mark it and continue
t.FailNow() - If test fails, then stop immidiately
t.Fatal() - log and fail now in one statement
*/
