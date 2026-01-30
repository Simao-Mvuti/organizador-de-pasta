package main

import (
	"fmt"
	"html/template"
	"net/http"
	"organizador/internal/organizador"
	"organizador/internal/scanner"
)

type PageData struct {
	Status string
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").Parse(`
		<h1>Organizador de Arquivos</h1>
		<form method="POST" action="/organizar">
			Pasta: <input type="text" name="pasta"><br><br>
			Modo: 
			<select name="modo">
				<option value="ext">Por Extensão</option>
				<option value="data">Por Data</option>
			</select><br><br>
			<input type="checkbox" name="dryrun"> Dry-Run (simulação)<br><br>
			<input type="submit" value="Organizar">
		</form>
		<p>{{.Status}}</p>
	`))
	tmpl.Execute(w, PageData{})
}

func organizarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pasta := r.FormValue("pasta")
	modo := r.FormValue("modo")
	dryRun := r.FormValue("dryrun") == "on"

	// Ler arquivos
	arquivos, err := scanner.LerPasta(pasta)
	if err != nil {
		tmpl := template.Must(template.New("home").Parse(`<p>Erro: {{.Status}}</p>`))
		tmpl.Execute(w, PageData{Status: err.Error()})
		return
	}

	// Organizar
	switch modo {
	case "data":
		organizador.OrganizarPorData(arquivos, pasta, dryRun)
	default:
		organizador.Organizar(arquivos, pasta, dryRun)
	}

	// Mensagem de sucesso
	tmpl := template.Must(template.New("home").Parse(`<p>Organização concluída com sucesso ✅</p><a href="/">Voltar</a>`))
	tmpl.Execute(w, PageData{Status: ""})
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/organizar", organizarHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
