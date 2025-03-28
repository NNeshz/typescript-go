package main

import (
	"strings"
)

type Token struct {
	Type  string // Ejemplo: "Keyword", "Identifier", "Type", "Operator", "Literal", etc.
	Value string // Ejemplo: "let", "x", "int", "+", "42", etc.
}

// tokenize -> Tomamos una cadena de texto como: "let x = 42;" y la convertimos en una lista de tokens.
func tokenize(code string) []Token {
	var tokens []Token
	var buffer strings.Builder // Para construir palabras carácter por carácter

	for _, char := range code {
		if char == ' ' || char == ';' || char == '=' || char == ':' {
			// Si hay algo en el buffer, lo convertimos a un token
			if buffer.Len() > 0 {
				value := buffer.String()
				var tokenType string
				switch value {
				case "let", "const", "var":
					tokenType = "Keyword"
				case "int", "string", "bool":
					tokenType = "Type"
				case "+", "-", "*", "/":
					tokenType = "Operator"
				default:
					tokenType = "Identifier"
				}

				tokens = append(tokens, Token{Type: tokenType, Value: value})
				buffer.Reset() // Limpiamos el buffer
			}

			// Arreglamos el simbolo como token sí no es un espacio
			if char != ' ' {
				tokens = append(tokens, Token{Type: "Punctuation", Value: string(char)})
			}
		} else {
			// Si el carácter no es un espacio, lo añadimos al buffer
			buffer.WriteRune(char)
		}
	}

	if buffer.Len() > 0 {
		value := buffer.String()
		var tokenType string
		switch value {
		case "let", "const":
			tokenType = "Keyword"
		case "int", "string", "bool":
			tokenType = "Type"
		case "+", "-", "*", "/":
			tokenType = "Operator"
		default:
			tokenType = "Identifier"
		}

		tokens = append(tokens, Token{Type: tokenType, Value: value})
	}

	return tokens
}

func main() {
	code := "let x: number = 10;"
	tokens := tokenize(code)

	for _, token := range tokens {
		println("Type:", token.Type, "Value:", token.Value)
	}
}
