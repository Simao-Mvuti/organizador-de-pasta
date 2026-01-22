package organizador

import (
	"fmt"
	"organizador/internal/backup"
	"organizador/internal/scanner"
	"os"
	"path/filepath"
	"strings"
)

// Organizador -- organiza a pasta com base na extensao do arquivo
func Organizar(arquivos []os.DirEntry, base string, dryRun bool) {
	for _, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(arquivo.Name()), ".")
		if ext == "" {
			ext = "outros"
		}

		pastaDestino := filepath.Join(base, strings.ToUpper(ext))
		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(pastaDestino, arquivo.Name())

		backupPath, err := backup.CriarBackup(origem, base)
		if err != nil {
			fmt.Println("Erro ao criar backup:", err)
			continue
		}
		fmt.Println("Backup criado:", backupPath)

		if scanner.Existe(destino) {
			fmt.Println("⚠️ Já existe:", destino)
			continue
		}

		if dryRun {
			fmt.Println("[SIMULAÇÃO]", origem, "->", destino)
			continue
		}

		if err := os.MkdirAll(pastaDestino, os.ModePerm); err != nil {
			fmt.Println("Erro ao criar pasta:", err)
			continue
		}

		if err := os.Rename(origem, destino); err != nil {
			fmt.Println("Erro ao mover arquivo:", err)
		}

	}
}

func OrganizarPorData(arquivos []os.DirEntry, base string, dryRun bool) {
	for _, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}

		info, err := arquivo.Info()
		if err != nil {
			fmt.Println("Erro ao ler info:", err)
			continue
		}

		data := info.ModTime()
		ano := data.Year()
		mes := int(data.Month())

		pastaDestino := filepath.Join(
			base,
			fmt.Sprintf("%d", ano),
			fmt.Sprintf("%02d", mes),
		)

		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(pastaDestino, arquivo.Name())

		if scanner.Existe(destino) {
			fmt.Println("⚠️ Arquivo já existe:", destino)
			continue
		}

		if dryRun {
			fmt.Println("[SIMULAÇÃO]", origem, "->", destino)
			continue
		}

		if err := os.MkdirAll(pastaDestino, os.ModePerm); err != nil {
			fmt.Println("Erro ao criar pasta:", err)
			continue
		}

		if err := os.Rename(origem, destino); err != nil {
			fmt.Println("Erro ao mover arquivo:", err)
		}
	}
}
