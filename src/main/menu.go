package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func menuAntigo(){


	fmt.Println("Criador de Organismos")


	var executando = true
	var opcao string = ""

	for (executando == true) {

		fmt.Println("Escolha uma opção [ cp : Criar Produtor, cc : Criar Consumidor, l: Listar, r: Remover s: Sair] :  ")
		fmt.Scan(&opcao)
		opcao = strings.ToLower(opcao)

		switch opcao {
		case "cp":
			ProdutorCriar()
		case "cc":
			ConsumidorCriar()
		case "l":
			OrganismoListar()
		case "r":
			OrganismoRemover()
		case "s":
			executando = false
		default:
			fmt.Println("Opção Desconhecida !!!")
		}
	}



}


func OrganismoListar() {



	fmt.Println("")


	fmt.Println("-------------- Listar Organismos --------------")

	fmt.Println("")

	files, err := ioutil.ReadDir(caminho)
	for _, file := range files {

		if strings.HasSuffix(file.Name(), ".organismo") {

			var nomearquivo string=strings.Replace(file.Name(),".organismo","",1)


			data, _ := ioutil.ReadFile(caminho + file.Name())

			note := &Organismo{}

			_ = xml.Unmarshal([]byte(data), &note)


			fmt.Println("    - " +  nomearquivo + "   :   " + note.Base.Tipo)

		}
	}
	err=err

	fmt.Println("")
	fmt.Println("")


}

func OrganismoRemover() {


	fmt.Println("")


	var nome string


	fmt.Println("-------------- Remover Organismo --------------")

	fmt.Println("")


	fmt.Print(" - Nome : ")
	fmt.Scan(&nome)


	os.Remove(caminho+ nome + ".organismo")

	fmt.Println("")
	fmt.Println("")

}

