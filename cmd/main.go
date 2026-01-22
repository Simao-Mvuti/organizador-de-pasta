package main

import (
	"fmt"
	"organizador/internal/organizador"
	"organizador/internal/scanner"
	"os"
)

func main() {
	var pasta string
	var err error

	pasta, err = os.Getwd()

	if err != nil {
		fmt.Println("Erro ao ler o diretório atual")
		return
	}

	if len(os.Args) > 1 {
		pasta = os.Args[1]
	}

	arquivo, err := scanner.LerPasta(pasta)

	if err != nil {
		fmt.Println("Erro", err)
		return
	}

	organizador.Organizar(arquivo, pasta, false)
	fmt.Println("Organizaçção Concluida com Sucesso")
}
