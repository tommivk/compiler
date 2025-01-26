package tokenizer

import (
	"reflect"
	"testing"
)

func SlicesEqual(t *testing.T, a, b []string) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v%s%v", a, " did not equal ", b)
	}
}

func TestShouldNotIncludeWhitespace(t *testing.T) {
	res, _ := Tokenize("  \n ")
	if len(res) > 0 {
		t.Fatalf("")
	}
}

func TestIntegersShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("2 25 9000")
	expected := []string{"2", "25", "9000"}
	SlicesEqual(t, res, expected)
}

func TestIfAndWhileKeywordsShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("if while while if")
	expected := []string{"if", "while", "while", "if"}
	SlicesEqual(t, res, expected)
}

func TestShouldWorkWithLargeAmountOfWhitespace(t *testing.T) {
	res, _ := Tokenize("if      \n while")
	expected := []string{"if", "while"}
	SlicesEqual(t, res, expected)
}

func TestKeywordsShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("a Variable _T b52")
	expected := []string{"a", "Variable", "_T", "b52"}
	SlicesEqual(t, res, expected)
}

func TestUnallowedKeywordShouldFail(t *testing.T) {
	_, err := Tokenize("53b")
	if err == nil {
		t.Fatalf("")
	}
}
