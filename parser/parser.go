package parser

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/tommivk/compiler/tokenizer"
)

type Expression interface{}

type Literal struct {
	Value int // int | bool?
}

type Identifier struct {
	Name string
}

type BinaryOp struct {
	Left  Expression
	Op    string
	Right Expression
}

func peek(pos int, tokens []tokenizer.Token) tokenizer.Token {
	if pos < len(tokens) {
		return tokens[pos]
	}
	return tokenizer.Token{Location: tokens[len(tokens)-1].Location, Type: "end"}
}

func consume(expected []string, pos *int, tokens []tokenizer.Token) tokenizer.Token {
	token := peek(*pos, tokens)
	if expected != nil && !slices.Contains(expected, token.Text) {
		log.Fatal("Token: ", token.Text, " Expected one of: ", strings.Join(expected, ", "))
	}
	*pos += 1
	return token
}

func parseIntLiteral(pos *int, tokens []tokenizer.Token) Literal {
	if token := peek(*pos, tokens); token.Type != "integer" {
		log.Fatal("Expected integer got: ", token.Type)
	}
	token := consume(nil, pos, tokens)
	value, err := strconv.Atoi(token.Text)
	if err != nil {
		log.Fatal("Failed to convert string", token.Text, " to integer")
	}
	return Literal{Value: value}
}

func parseIdentifier(pos *int, tokens []tokenizer.Token) Identifier {
	if token := peek(*pos, tokens); token.Type != "identifier" {
		log.Fatal("Expected identifier got: ", token.Type)
	}
	token := consume(nil, pos, tokens)
	return Identifier{Name: token.Text}
}

func parseTerm(pos *int, tokens []tokenizer.Token) Expression {
	left := parseFactor(pos, tokens)
	for slices.Contains([]string{"*", "/"}, peek(*pos, tokens).Text) {
		operator := consume(nil, pos, tokens)
		right := parseFactor(pos, tokens)
		left = BinaryOp{Left: left, Op: operator.Text, Right: right}
	}
	return left
}

func parseParenthized(pos *int, tokens []tokenizer.Token) Expression {
	consume([]string{"("}, pos, tokens)
	expression := parseExpression(pos, tokens)
	consume([]string{")"}, pos, tokens)
	return expression
}

func parseFactor(pos *int, tokens []tokenizer.Token) Expression {
	if peek(*pos, tokens).Text == "(" {
		return parseParenthized(pos, tokens)
	}
	if peek(*pos, tokens).Type == "integer" {
		return parseIntLiteral(pos, tokens)
	}
	if peek(*pos, tokens).Type == "identifier" {
		return parseIdentifier(pos, tokens)
	}
	log.Fatal(peek(*pos, tokens).Location, " Expected integer or identifier, got: ", peek(*pos, tokens).Type)
	return nil
}

func parseExpression(pos *int, tokens []tokenizer.Token) Expression {
	left := parseTerm(pos, tokens)
	for slices.Contains([]string{"+", "-"}, peek(*pos, tokens).Text) {
		operator := consume(nil, pos, tokens)
		right := parseTerm(pos, tokens)
		left = BinaryOp{Left: left, Op: operator.Text, Right: right}
	}
	return left
}

func Parse(tokens []tokenizer.Token) Expression {
	pos := 0
	return parseExpression(&pos, tokens)
}
