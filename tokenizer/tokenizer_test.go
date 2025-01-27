package tokenizer

import (
	"fmt"
	"reflect"
	"testing"
)

func SlicesEqual(t *testing.T, a, b []string) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v%s%v", a, " did not equal ", b)
	}
}

func TokensEqual(t *testing.T, a, b []Token) {
	if len(a) != len(b) {
		t.Fatalf("%v%s%v", a, " did not equal ", b)
	}
	for i, v := range a {
		if v.Text != b[i].Text || v.Type != b[i].Type {
			t.Fatalf("%v%s%v", v, " did not equal ", b[i])
		}
	}
}

func LocationsEqual(t *testing.T, a, b []Token) {
	if len(a) != len(b) {
		t.Fatalf("%v%s%v", a, " did not equal ", b)
	}
	for i, v := range a {
		if v.Location != b[i].Location {
			t.Fatalf("%v%s%v", v.Location, " did not equal ", b[i].Location)
		}
	}
}

func TestShouldNotIncludeWhitespace(t *testing.T) {
	res, _ := Tokenize("  \n ")
	if len(res) > 0 {
		t.Fatalf("")
	}
}

func TestIntegersShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("2 25 9000 a123a")
	expected := []Token{
		{Text: "2", Type: "integer"},
		{Text: "25", Type: "integer"},
		{Text: "9000", Type: "integer"},
		{Text: "a123a", Type: "identifier"}}
	TokensEqual(t, res, expected)
}

func TestIfAndWhileKeywordsShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("if while while if")
	expected := []Token{
		{Text: "if", Type: "identifier"},
		{Text: "while", Type: "identifier"},
		{Text: "while", Type: "identifier"},
		{Text: "if", Type: "identifier"}}
	TokensEqual(t, res, expected)
}

func TestShouldWorkWithLargeAmountOfWhitespace(t *testing.T) {
	res, _ := Tokenize("if      \n while")
	expected := []Token{
		{Text: "if", Type: "identifier"},
		{Text: "while", Type: "identifier"}}
	TokensEqual(t, res, expected)
}

func TestKeywordsShouldBeIncluded(t *testing.T) {
	res, _ := Tokenize("a Variable _T b52")
	expected := []Token{
		{Text: "a", Type: "identifier"},
		{Text: "Variable", Type: "identifier"},
		{Text: "_T", Type: "identifier"},
		{Text: "b52", Type: "identifier"},
	}
	TokensEqual(t, res, expected)
}

func TestLocations(t *testing.T) {
	res, _ := Tokenize("1   2\n3  \n\n\n    4")
	fmt.Println(res)
	expected := []Token{
		{Text: "1", Type: "integer", Location: Location{Line: 0, Column: 0}},
		{Text: "2", Type: "integer", Location: Location{Line: 0, Column: 4}},
		{Text: "3", Type: "integer", Location: Location{Line: 1, Column: 0}},
		{Text: "4", Type: "integer", Location: Location{Line: 4, Column: 4}},
	}
	LocationsEqual(t, res, expected)
}

func TestOperators(t *testing.T) {
	res, _ := Tokenize(">=<=-=*==<>!=/+")
	expected := []Token{
		{Text: ">=", Type: "operator"},
		{Text: "<=", Type: "operator"},
		{Text: "-", Type: "operator"},
		{Text: "=", Type: "operator"},
		{Text: "*", Type: "operator"},
		{Text: "==", Type: "operator"},
		{Text: "<", Type: "operator"},
		{Text: ">", Type: "operator"},
		{Text: "!=", Type: "operator"},
		{Text: "/", Type: "operator"},
		{Text: "+", Type: "operator"},
	}
	TokensEqual(t, res, expected)
}

func TestPunctuations(t *testing.T) {
	res, _ := Tokenize(")({},;")
	expected := []Token{
		{Text: ")", Type: "punctuation"},
		{Text: "(", Type: "punctuation"},
		{Text: "{", Type: "punctuation"},
		{Text: "}", Type: "punctuation"},
		{Text: ",", Type: "punctuation"},
		{Text: ";", Type: "punctuation"},
	}
	TokensEqual(t, res, expected)
}

func TestFunctionCall(t *testing.T) {
	res, _ := Tokenize("var a = test(input)")
	expected := []Token{
		{Text: "var", Type: "identifier"},
		{Text: "a", Type: "identifier"},
		{Text: "=", Type: "operator"},
		{Text: "test", Type: "identifier"},
		{Text: "(", Type: "punctuation"},
		{Text: "input", Type: "identifier"},
		{Text: ")", Type: "punctuation"},
	}
	TokensEqual(t, res, expected)
}

func TestUnallowedKeywordShouldFail(t *testing.T) {
	_, err := Tokenize("53b")
	if err == nil {
		t.Fatalf("")
	}
}
