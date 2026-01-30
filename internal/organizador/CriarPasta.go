package organizador

import "os"

func CriarPasta(caminho string) error {
	return os.MkdirAll(caminho, os.ModePerm)
}

func MoverArquivo(origem, destino string) error {
	return os.Rename(origem, destino)
}
