package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)



func ProdutorCriar() {

	fmt.Println("")


	var nome string
	var adulto int
	var vida int
	var reproducao_frequencia int
var cor uint32

	fmt.Println("-------------- Criar Produtor --------------")

	fmt.Println("")


	fmt.Print(" - Nome : ")
	fmt.Scan(&nome)

	fmt.Print(" - Cor : ")
	fmt.Scan(&cor)

	fmt.Print(" - Tempo Adulto : ")
	fmt.Scan(&adulto)

	fmt.Print(" - Tempo Vida : ")
	fmt.Scan(&vida)

	fmt.Print(" - Frequencia de Reprodução : ")
	fmt.Scan(&reproducao_frequencia)



	x := &Organismo{
		Base: Base{Tipo: "Produtor", Adulto: adulto, Vida: vida,Cor:cor},
		Reproducao:Reproducao{Frequencia:reproducao_frequencia},
	}

	file, _ := os.Create( caminho + nome+".organismo" )
	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("", "  ")
	if err := enc.Encode(x); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("")
	fmt.Println("")

}


