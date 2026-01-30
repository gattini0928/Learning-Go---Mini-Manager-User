package main

import ("fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Bem-vindo Luna! Ao sistema de registro de usuários.")
	users := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Por favor selecione uma opção:")
		fmt.Println("1. Registrar novo usuário")
		fmt.Println("2. Mostrar usuário")
		fmt.Println("3. Deletar usuário")
		fmt.Println("4. Sair")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Tente Novamente, opção inválida.")
			continue
		}

		if choice == 1 {
			fmt.Print("Nome Completo: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			fmt.Print("Senha: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)
			var ok bool
			users, ok = registerUser(name, email, password, users)
			if !ok {
				continue
			}

		} else if choice == 2 {
			fmt.Print("Nome Completo: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			showUser(name, email, users)
		} else if choice == 3 {
			fmt.Print("Digite o email do usuário que quer deletar: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			users = deleteUser(email, users)
		} else if choice == 4 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Tente Novamente, opção inválida.")
		}
	}
}
