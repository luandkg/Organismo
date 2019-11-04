package main

import (
	"io/ioutil"
	"strings"
)

func listar(local string) []string {

	var conteudo []string


	files, err := ioutil.ReadDir(local)
	for _, file := range files {

		if strings.HasSuffix(file.Name(), ".organismo") {

			var nomearquivo string=strings.Replace(file.Name(),".organismo","",1)
			conteudo = append(conteudo, nomearquivo)


			//data, _ := ioutil.ReadFile(caminho + file.Name())

			//note := &Organismo{}

			//_ = xml.Unmarshal([]byte(data), &note)

		}
	}
	err=err

return conteudo
}
