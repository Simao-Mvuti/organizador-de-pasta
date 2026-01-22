package organizador

import (
	"organizador/internal/scanner"
	"os"
	"path/filepath"
	"testing"
)

// Teste Organizar por extensão
func TestOrganizar(t *testing.T) {
	// 1️⃣ Criar pasta temporária
	base := t.TempDir() // pasta temporária que Go limpa automaticamente

	// 2️⃣ Criar arquivos de teste
	os.WriteFile(filepath.Join(base, "teste1.txt"), []byte("ok"), 0644)
	os.WriteFile(filepath.Join(base, "teste2.pdf"), []byte("ok"), 0644)
	os.WriteFile(filepath.Join(base, "teste3"), []byte("ok"), 0644) // sem extensão

	// 3️⃣ Ler pasta
	arquivos, err := scanner.LerPasta(base)
	if err != nil {
		t.Fatal("Erro ao ler pasta:", err)
	}

	// 4️⃣ Rodar organizador
	Organizar(arquivos, base, false)

	// 5️⃣ Verificar resultados
	if _, err := os.Stat(filepath.Join(base, "TXT", "teste1.txt")); err != nil {
		t.Error("Arquivo teste1.txt não movido corretamente")
	}

	if _, err := os.Stat(filepath.Join(base, "PDF", "teste2.pdf")); err != nil {
		t.Error("Arquivo teste2.pdf não movido corretamente")
	}

	if _, err := os.Stat(filepath.Join(base, "OUTROS", "teste3")); err != nil {
		t.Error("Arquivo teste3 não movido corretamente")
	}
}

func TestOrganizarPorData(t *testing.T) {
	base := t.TempDir()

	// 1️⃣ Criar arquivo
	arquivoPath := filepath.Join(base, "arquivo.txt")
	os.WriteFile(arquivoPath, []byte("ok"), 0644)

	arquivos, _ := scanner.LerPasta(base)

	// 2️⃣ Rodar organização
	OrganizarPorData(arquivos, base, false)

	// 3️⃣ Procurar arquivo na pasta de destino (ano/mes)
	var encontrado bool
	filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.Name() == "arquivo.txt" {
			encontrado = true
		}
		return nil
	})

	if !encontrado {
		t.Error("Arquivo não movido corretamente por data")
	}
}
