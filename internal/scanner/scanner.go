package scanner

import (
	"os"
	"path/filepath"
)

// LerPasta -- retorna uma lista de diretorio
func LerPasta(pasta string) ([]os.DirEntry, error) {
	return os.ReadDir(filepath.Clean(pasta))

}
