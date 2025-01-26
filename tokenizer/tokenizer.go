package tokenizer

import (
	"errors"
	"regexp"
)

func CheckWhiteSpace(s string) int {
	re := regexp.MustCompile(`(^\s*)`)
	result := re.Find([]byte(s))
	whitespace := len(string(result))
	return whitespace
}

func ExtractNextToken(s string) string {
	re := regexp.MustCompile(`[^\s]+`)
	token := re.Find([]byte(s))
	return string(token)
}

func MatchRegexes(s string) error {
	// Check integers
	re := regexp.MustCompile(`^[0-9]*$`)
	result := re.Find([]byte(s))
	integer := string(result)
	if len(integer) > 0 {
		return nil
	}

	// Check variable names and keywords `[^\s]+`
	re = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z_0-9]*`)
	result = re.Find([]byte(s))
	word := string(result)
	if len(word) > 0 {
		return nil
	}

	return errors.New("error while matching regex")
}

func Tokenize(s string) ([]string, error) {
	tokens := []string{}
	for i := 0; i < len(s); i++ {
		whitespace := CheckWhiteSpace(s[i:])
		if whitespace > 0 {
			i += whitespace - 1
			continue
		}

		token := ExtractNextToken(s[i:])
		err := MatchRegexes(token)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
		i += len(token)
	}
	return tokens, nil
}
