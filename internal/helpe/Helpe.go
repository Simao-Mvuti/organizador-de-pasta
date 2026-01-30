package helpe

import "fmt"

func MostrarHelp() {
	fmt.Println("Organizador de arquivos ✅")
	fmt.Println()
	fmt.Println("Uso:")
	fmt.Println("  organizador [OPÇÕES] [PASTA]")
	fmt.Println()
	fmt.Println("Opções:")
	fmt.Println("  --exec       Organizar por extensão (padrão)")
	fmt.Println("  --data      Organizar por data (ANO/MÊS)")
	fmt.Println("  --teste  Simula a organização sem mover arquivos")
	fmt.Println("  --ajuda      Mostra esta mensagem")
	fmt.Println()
	fmt.Println("Exemplos:")
	fmt.Println("  ./organizador /home/samuel/arquivos")
	fmt.Println("  ./organizador --data --teste /home/samuel/arquivos")
	fmt.Println("  ./organizador --exec /home/samuel/arquivos")
}
