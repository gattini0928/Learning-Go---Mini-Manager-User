package main

import ("fmt"
	"strings"
	"unicode"
)

func validateName(name string) (string , bool) {
	parts := strings.Fields(name)
	if name == "" {
		fmt.Println("O nome não pode estar vazio.")
		return name, false
	}
	if len(parts) < 2 {
		fmt.Println("Por favor escreva o seu nome e sobrenome.")
		return name, false
	}
	firstName := parts[0]
	lastName := parts[1]
	if len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Por favor escreva um nome e sobrenome com pelo menos 2 caracteres.")
		return name, false
	}
	return name, true
}

func validateEmail(email string) (string, bool) {
	if email == "" {
		fmt.Println("O email não pode estar vazio.")
		return email, false
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		fmt.Println("Por favor escreva um email válido.")
		return email, false
	}

	if strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") {
		fmt.Println("Por favor escreva um email válido.")
		return email, false
	}

	return email, true
}

func validatePassword(password string) (string , bool) {
	if password == "" {
		fmt.Println("A senha não pode estar vazia.")
		return password, false
	}

	if len(password) < 6 {
		fmt.Println("A senha deve ter pelo menos 6 caracteres.")
		return password, false
	}

	digits := 0
	letters:= 0
	specials := 0
	for _, char := range password {
		if unicode.IsDigit(char) {
			digits++
		} else if unicode.IsLetter(char) {
			letters++
		} else {
			specials++
		}
	}
	if digits > 0 && letters > 0 && specials > 0 {
		return password, true
	} else {
		fmt.Println("A senha deve conter pelo menos uma letra, um número e um caractere especial.")
		return password, false
	}
}