package main

import "fmt"

func registerUser(name string, email string, password string, users []string) ([]string, bool){
	validName, okName := validateName(name)
	validEmail, okEmail := validateEmail(email)
	validPassword, okPassword := validatePassword(password)
	if !okName {
		return users, false
	}
	if !okEmail {
		return users, false
	}
	if !okPassword {
		return users, false
	}
	users = append(users, validName, validEmail, validPassword)
	fmt.Printf("Usuário(a) %s registrado com sucesso!\n", name)
	return users, true
}

func showUser(name string, email string, users []string) {
	for i := 0; i < len(users); i+=3 {
		if users[i] == name && users[i+1] == email {
			fmt.Printf("Usuário encontrado:\nNome: %s\nEmail: %s\n", users[i], users[i+1])
			return
		}
	}
	fmt.Println("Usuário não encontrado.")
}

func deleteUser(email string, users[]string) []string {
	for i := 0; i < len(users); i+=3 {
		if users[i+1] == email {
			users = append(users[:i], users[i+3:]...)
			fmt.Println("Usuário deletado com sucesso.")
			return users
		}
	}
	fmt.Println("Usuário não encontrado.")
	return users
}