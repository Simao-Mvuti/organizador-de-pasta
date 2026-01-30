package main

import (
	"fmt"
	"organizador/internal/helpe"
	"organizador/internal/organizador"
	"organizador/internal/scanner"
	"organizador/util"
	"os"
)

func main() {
	// valores padrão
	pasta, err := os.Getwd()
	if err != nil {
		fmt.Println(util.Red, "Erro ao obter diretório atual", util.Reset)
		return
	}

	modo := "ext"
	dryRun := false

	// leitura dos argumentos
	for _, arg := range os.Args[1:] {
		switch arg {
		case "ajuda":
			helpe.MostrarHelp()
		case "--data":
			modo = "data"
		case "--exec":
			modo = "ext"
		case "--teste":
			dryRun = true
		default:
			// se não começa com --, assume que é pasta
			if len(arg) > 2 && arg[:2] != "--" {
				pasta = arg
			}
		}
	}

	arquivo, err := scanner.LerPasta(pasta)
	if err != nil {
		fmt.Println(util.Red, "Erro:", err, util.Reset)
		return
	}

	switch modo {
	case "data":
		organizador.OrganizarPorData(arquivo, pasta, dryRun)
	default:
		organizador.Organizar(arquivo, pasta, dryRun)
	}

	fmt.Println(util.Green, "Organização concluída com sucesso ", util.Reset)

}
