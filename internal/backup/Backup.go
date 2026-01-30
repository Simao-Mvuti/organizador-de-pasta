package backup

import (
	"fmt"
	"organizador/internal/scanner"
	"os"
	"path/filepath"
	"strings"
)

func CriarBackup(origem string, base string) (string, error) {
	backupBase := filepath.Join(base, "backup")

	// cria pasta backup se não existir
	if err := os.MkdirAll(backupBase, os.ModePerm); err != nil {
		return "", err
	}

	// gera caminho único para backup
	destino := filepath.Join(backupBase, filepath.Base(origem))
	contador := 1
	for scanner.Existe(destino) {
		nome := strings.TrimSuffix(filepath.Base(origem), filepath.Ext(origem))
		ext := filepath.Ext(origem)
		destino = filepath.Join(backupBase, fmt.Sprintf("%s (%d)%s", nome, contador, ext))
		contador++
	}

	// copia o arquivo
	input, err := os.ReadFile(origem)
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(destino, input, 0644); err != nil {
		return "", err
	}

	return destino, nil
}
