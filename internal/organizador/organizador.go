package organizador

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Organizador -- organiza a pasta com base na extensao do arquivo
func Organizar(arquivos []os.DirEntry, base string) {
	for _, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(arquivo.Name()), ".")
		if ext == "" {
			ext = "outros"
		}

		pastaDestino := filepath.Join(base, strings.ToUpper(ext))

		if err := os.MkdirAll(pastaDestino, os.ModePerm); err != nil {
			fmt.Println("Erro ao criar pasta:", err)
			continue
		}

		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(pastaDestino, arquivo.Name())

		if err := os.Rename(origem, destino); err != nil {
			fmt.Println("Erro ao mover arquivo:", err)
		}
	}
}
