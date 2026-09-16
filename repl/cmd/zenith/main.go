package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jvbenetti/zenith-lang.git/repl"
)

func main() {
	// Get user from OS
	userSys, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Mensagem de boas-vindas customizada da Zenith!
	fmt.Printf("Olá %s! Este é o terminal da linguagem Zenith!\n", userSys.Username)
	fmt.Printf("Sinta-se livre para digitar alguns comandos\n")

	// Inicia o terminal passando a entrada do teclado (os.Stdin) e a tela (os.Stdout)
	repl.Start(os.Stdin, os.Stdout)
}
