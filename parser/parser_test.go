package parser

import (
	"fmt"
	"testing"

	"github.com/tommivk/compiler/tokenizer"
)

func TestShouldParseSimplePlusOperation(t *testing.T) {
	input := []tokenizer.Token{{Text: "1", Type: "integer"}, {Text: "+", Type: "operator"}, {Text: "2", Type: "integer"}}
	result := Parse(input)
	fmt.Println(result)
}
