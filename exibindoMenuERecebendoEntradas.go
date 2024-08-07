package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string
	var matricula int
	var comando int
	versao := 1.3
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Print("Agora, sua matrícula: ")
	fmt.Scan(&matricula)
	fmt.Println()
	fmt.Println("Bem vindo, sr.", nome)
	fmt.Println("id:", matricula, reflect.TypeOf(matricula))
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
	fmt.Println("Diga-me, o que deseja fazer?")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
	fmt.Print(">>> ")
	fmt.Scan(&comando)
	fmt.Println()
	fmt.Println("-> A opção escolhida foi", comando, "e o seu endereço na memória é:", &comando)
}

// Agora, preciso aprender como tratar e corrigir erros.
