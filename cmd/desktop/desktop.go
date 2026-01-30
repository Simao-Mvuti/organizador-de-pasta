package main

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"organizador/internal/organizador"
	"organizador/internal/scanner"
)

// main é o ponto de entrada do programa desktop
func main() {
	// Cria a aplicação Fyne
	a := app.New()

	// Cria uma janela com título
	w := a.NewWindow("Organizador de Arquivos")

	// Label que mostra mensagens de status/logs
	status := widget.NewMultiLineEntry()
	status.SetPlaceHolder("Logs aparecerão aqui...")

	// Entrada de texto para o usuário informar a pasta
	pastaEntry := widget.NewEntry()
	pastaEntry.SetPlaceHolder("Digite ou cole o caminho da pasta")

	// Checkbox para Dry-Run (simulação)
	dryRunCheck := widget.NewCheck("Dry-Run (simulação)", func(bool) {})

	// Função para processar a organização, evita repetir código
	processar := func(modo string) {
		pasta := pastaEntry.Text
		if pasta == "" {
			status.SetText("⚠️ Por favor, informe a pasta!")
			return
		}

		// Ler arquivos na pasta
		arquivos, err := scanner.LerPasta(pasta)
		if err != nil {
			status.SetText(fmt.Sprintf("❌ Erro ao ler a pasta: %v", err))
			return
		}

		// Limpar logs antes de começar
		status.SetText("Iniciando organização...\n")

		// Função para atualizar logs na interface
		log := func(msg string) {
			status.SetText(status.Text + msg + "\n")
		}

		// Escolher modo de organização
		switch modo {
		case "ext":
			// Organiza por extensão, passando dry-run
			for _, arquivo := range arquivos {
				if arquivo.IsDir() {
					continue
				}

				ext := filepath.Ext(arquivo.Name())
				if ext == "" {
					ext = "outros"
				}

				pastaDestino := filepath.Join(pasta, ext)
				origem := filepath.Join(pasta, arquivo.Name())
				destino := filepath.Join(pastaDestino, arquivo.Name())

				// Dry-run
				if dryRunCheck.Checked {
					log(fmt.Sprintf("[SIMULAÇÃO] %s -> %s", origem, destino))
					continue
				}

				// Criar pasta se não existir
				if err := organizador.CriarPasta(pastaDestino); err != nil {
					log(fmt.Sprintf("❌ Erro ao criar pasta: %v", err))
					continue
				}

				// Mover arquivo
				if err := organizador.MoverArquivo(origem, destino); err != nil {
					log(fmt.Sprintf("❌ Erro ao mover arquivo: %v", err))
					continue
				}

				log(fmt.Sprintf("✅ [MOVIDO] %s -> %s", origem, destino))
			}

		case "data":
			// Organiza por data, dry-run
			organizador.OrganizarComLog(arquivos, pasta, dryRunCheck.Checked, log)
		}

		log("🎉 Organização concluída!")
	}

	// Botões para escolher o modo de organização
	btnExt := widget.NewButton("Organizar por Extensão", func() { processar("ext") })
	btnData := widget.NewButton("Organizar por Data", func() { processar("data") })

	// Organiza elementos na janela (vertical)
	content := container.NewVBox(
		widget.NewLabel("📂 Caminho da Pasta:"),
		pastaEntry,
		dryRunCheck,
		btnExt,
		btnData,
		widget.NewLabel("📄 Logs:"),
		status,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(500, 400)) // Define tamanho inicial da janela
	w.ShowAndRun()
}
