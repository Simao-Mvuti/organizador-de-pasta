package organizador

import (
	"fmt"
	"organizador/internal/backup"
	"organizador/internal/scanner"
	"organizador/util"
	"os"
	"path/filepath"
	"strings"
)

// Organizador -- organiza a pasta com base na extensao do arquivo
func Organizar(arquivos []os.DirEntry, base string, dryRun bool) {
	total := len(arquivos)

	for i, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}

		ext := strings.TrimPrefix(filepath.Ext(arquivo.Name()), "")
		if ext == "" {
			ext = "outros"
		}

		pastaDestino := filepath.Join(base, strings.ToUpper(ext))
		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(pastaDestino, arquivo.Name())

		// Backup
		backupPath, err := backup.CriarBackup(origem, base)
		if err != nil {
			fmt.Println(util.Red, "❌ Erro ao criar backup:", err, util.Reset)
			continue
		}
		fmt.Println(util.Yellow, "📦 Backup criado:", backupPath, util.Reset)

		// Verifica se já existe
		if scanner.Existe(destino) {
			fmt.Println(util.Yellow, "⚠️ Já existe:", destino, util.Reset)
			continue
		}

		// Dry-run
		if dryRun {
			fmt.Println(util.Blue, "[SIMULAÇÃO]", origem, "->", destino, util.Reset)
		} else {
			// Cria pasta se não existir
			if err := os.MkdirAll(pastaDestino, os.ModePerm); err != nil {
				fmt.Println(util.Red, "❌ Erro ao criar pasta:", err, util.Reset)
				continue
			}

			// Move arquivo
			if err := os.Rename(origem, destino); err != nil {
				fmt.Println(util.Red, "❌ Erro ao mover arquivo:", err, util.Reset)
				continue
			}
			fmt.Println(util.Green, "[MOVIDO]", origem, "->", destino, util.Reset)
		}

		// Barra de progresso
		fmt.Printf("\rProgresso: %d/%d arquivos processados", i+1, total)
	}

	fmt.Println("\n" + util.Green + "✅ Organização concluída!" + util.Reset)
}

func OrganizarPorData(arquivos []os.DirEntry, base string, dryRun bool) {
	total := len(arquivos)

	for i, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}

		info, err := arquivo.Info()
		if err != nil {
			fmt.Println(util.Red, "❌ Erro ao ler info:", err, util.Reset)
			continue
		}

		data := info.ModTime()
		ano := data.Year()
		mes := int(data.Month())

		pastaDestino := filepath.Join(base, fmt.Sprintf("%d", ano), fmt.Sprintf("%02d", mes))
		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(pastaDestino, arquivo.Name())

		// Backup
		backupPath, err := backup.CriarBackup(origem, base)
		if err != nil {
			fmt.Println(util.Red, "❌ Erro ao criar backup:", err, util.Reset)
			continue
		}
		fmt.Println(util.Yellow, "📦 Backup criado:", backupPath, util.Reset)

		// Verifica se já existe
		if scanner.Existe(destino) {
			fmt.Println(util.Yellow, "⚠️ Arquivo já existe:", destino, util.Reset)
			continue
		}

		// Dry-run
		if dryRun {
			fmt.Println(util.Blue, "[SIMULAÇÃO]", origem, "->", destino, util.Reset)
		} else {
			// Cria pasta e move arquivo
			if err := os.MkdirAll(pastaDestino, os.ModePerm); err != nil {
				fmt.Println(util.Red, "❌ Erro ao criar pasta:", err, util.Reset)
				continue
			}

			if err := os.Rename(origem, destino); err != nil {
				fmt.Println(util.Red, "❌ Erro ao mover arquivo:", err, util.Reset)
				continue
			}
			fmt.Println(util.Green, "[MOVIDO]", origem, "->", destino, util.Reset)
		}

		// Barra de progresso
		fmt.Printf("\rProgresso: %d/%d arquivos processados", i+1, total)
	}

	fmt.Println("\n" + util.Green + "✅ Organização concluída!" + util.Reset)
}
