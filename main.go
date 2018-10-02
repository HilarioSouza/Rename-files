package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//Declaração do array de arquivos
	var files []string
	//Pasta dos arquivos que serão listados (parâmetro passado na execução do  sistema)
	root, erro := getArg(1)
	if erro != nil {
		panic(erro)
	}
	//População do array de strings com os caminhos dos arquivos
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	//Se o objeto Erro não for nulo, então lance a exceção
	if err != nil {
		panic(err)
	}
	//Laço no array de arquivos
	for index, file := range files {
		//O primeiro item do laço é a pasta, portanto não me interessa. Não sei como iniciar o laço pelo índice 1
		if index > 0 {
			//nome do arquivo original
			before := file
			//nome do novo arquivo - Preciso melhorar a lógica da cópia do nome para tratar somente o nome do arquivo e depois concatenar a pasta e o nome
			path := filepath.Dir(file)
			name := filepath.Base(file)
			after := path + string(os.PathSeparator) + name[12:14] + ".csv"
			//Imprimir o nome dos arquivos
			fmt.Printf(before + " : " + after + "\n")
			//Renomear o arquivo
			os.Rename(before, after)
		}
	}
	//Imprimir mensagem de fim do processamento
	fmt.Println("DONE")
	//Finalizar a execução com status de sucesso.
	os.Exit(0)
}

func getArg(index int) (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("Nenhum parâmetro informado")
	}
	return os.Args[index], nil
}
