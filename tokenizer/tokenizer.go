package tokenizer

import (
	"errors"
	"regexp"
)

type Location struct {
	Line   int
	Column int
}

type Token struct {
	Text string
	Type string
	Location
}

func CountWhitespace(s string) int {
	re := regexp.MustCompile(`^[^\S\r\n]`)
	result := re.Find([]byte(s))
	whitespace := len(string(result))
	return whitespace
}

func CountNewlines(s string) int {
	re := regexp.MustCompile(`^[\r\n]+`)
	result := re.Find([]byte(s))
	newLines := len(string(result))
	return newLines
}

func ExtractNextToken(s string) string {
	re := regexp.MustCompile(`[^\s]+`)
	token := re.Find([]byte(s))
	return string(token)
}

func MatchInteger(s string) string {
	re := regexp.MustCompile(`^[0-9]*$`)
	result := re.Find([]byte(s))
	integer := string(result)
	return integer
}

func MatchIdentifiers(s string) string {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z_0-9]*`)
	result := re.Find([]byte(s))
	word := string(result)
	return word
}

func MatchOperators(s string) string {
	re := regexp.MustCompile(`(\+|-|\*|\/|==|=|!=|<=|>=|<|>)`)
	result := re.Find([]byte(s))
	operator := string(result)
	return operator
}

func MatchPunctuation(s string) string {
	re := regexp.MustCompile(`(\(|,|;|}|{|\))`)
	result := re.Find([]byte(s))
	punctuation := string(result)
	return punctuation
}

func MatchRegexes(s string) (Token, error) {
	if integer := MatchInteger(s); integer != "" {
		return Token{Type: "integer", Text: integer}, nil
	}

	if identifier := MatchIdentifiers(s); identifier != "" {
		return Token{Type: "identifier", Text: identifier}, nil
	}
	if operator := MatchOperators(s); operator != "" {
		return Token{Type: "operator", Text: operator}, nil
	}

	if punctuation := MatchPunctuation(s); punctuation != "" {
		return Token{Type: "punctuation", Text: punctuation}, nil
	}

	return Token{}, errors.New("error while matching regex")
}

var line int
var column int

func Tokenize(s string) ([]Token, error) {
	tokens := []Token{}
	line = 0
	column = 0
	for i := 0; i < len(s); i++ {
		newLines := CountNewlines(s[i:])
		if newLines > 0 {
			line += newLines
			column = 0
			i += newLines - 1
			continue
		}

		whitespace := CountWhitespace(s[i:])
		if whitespace > 0 {
			column += whitespace
			i += whitespace - 1
			continue
		}

		input := ExtractNextToken(s[i:])
		token, err := MatchRegexes(input)
		if err != nil {
			return nil, err
		}
		token.Location = Location{Line: line, Column: column}
		tokens = append(tokens, token)
		i += len(token.Text) - 1
		column += len(token.Text)
	}

	return tokens, nil
}
