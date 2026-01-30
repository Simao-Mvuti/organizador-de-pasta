package organizador

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizarComLog(arquivos []os.DirEntry, base string, dryRun bool, log func(string)) {
	for _, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}
		// exemplo simplificado
		origem := filepath.Join(base, arquivo.Name())
		destino := filepath.Join(base, "outros", arquivo.Name())

		if dryRun {
			log(fmt.Sprintf("[SIMULAÇÃO] %s -> %s", origem, destino))
			continue
		}

		err := os.MkdirAll(filepath.Dir(destino), os.ModePerm)
		if err != nil {
			log(fmt.Sprintf("❌ Erro ao criar pasta: %v", err))
			continue
		}

		err = os.Rename(origem, destino)
		if err != nil {
			log(fmt.Sprintf("❌ Erro ao mover arquivo: %v", err))
			continue
		}

		log(fmt.Sprintf("✅ [MOVIDO] %s -> %s", origem, destino))
	}
}
