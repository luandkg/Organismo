package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"io"
	"os"
	"strconv"
)

func consumidores(boxPrincipal* gtk.Box){


	boxTituloProdutor,_:= gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	boxMestreProdutor ,_:=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)
	boxSalvadorProdutor,_:= gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)

	boxPrincipal.Add(boxTituloProdutor)
	boxPrincipal.Add(boxMestreProdutor)
	boxPrincipal.Add(boxSalvadorProdutor)


	labelTituloProdutor,_ := gtk.LabelNew("Criar Consumidor")
	boxTituloProdutor.PackStart(labelTituloProdutor, true, true, 0)

	boxTitulos,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,5)
	boxEntradas ,_:= gtk.BoxNew(gtk.ORIENTATION_VERTICAL,5)

	boxMestreProdutor.PackStart(boxTitulos, true, true, 0)
	boxMestreProdutor.PackStart(boxEntradas, true, true, 0)


	// Define os Rotulos

	labelNome,_ := gtk.LabelNew("Nome")
	labelCor,_ := gtk.LabelNew("Cor")
	labelAdulto,_ := gtk.LabelNew("Tempo Adulto")
	labelVida,_ := gtk.LabelNew("Tempo Vida")
	labelFreReproducao,_ := gtk.LabelNew("Frequência de Reprodução")
	labelGestacao,_ := gtk.LabelNew("Tempo de Gestação")
	labelGasCarbonico,_ := gtk.LabelNew("Gás Carbônico")
	labelGasOxigenio,_ := gtk.LabelNew("Gás Oxigênio")


	// Define as Entradas

	textoNome,_:=gtk.EntryNew()
	textoCor,_:=gtk.EntryNew()
	textoAdulto,_:=gtk.EntryNew()
	textoVida,_:=gtk.EntryNew()
	textoFreReproducao,_:=gtk.EntryNew()
	textoGestacao,_:=gtk.EntryNew()
	textoGasCarbonico,_:=gtk.EntryNew()
	textoGasOxigenio,_:=gtk.EntryNew()

	// Define os valores iniciais

	textoNome.SetText("Consumidor")
	textoCor.SetText("1618688")
	textoAdulto.SetText("100")
	textoVida.SetText("10000")
	textoFreReproducao.SetText("100")
	textoGestacao.SetText("50")
	textoGasCarbonico.SetText("1")
	textoGasOxigenio.SetText("1")



	boxTitulos.PackStart(labelNome, true, true, 0)
	boxEntradas.PackStart(textoNome, true, true, 0)

	boxTitulos.PackStart(labelCor, true, true, 0)
	boxEntradas.PackStart(textoCor, true, true, 0)

	boxTitulos.PackStart(labelAdulto, true, true, 0)
	boxEntradas.PackStart(textoAdulto, true, true, 0)

	boxTitulos.PackStart(labelVida, true, true, 0)
	boxEntradas.PackStart(textoVida, true, true, 0)

	boxTitulos.PackStart(labelFreReproducao, true, true, 0)
	boxEntradas.PackStart(textoFreReproducao, true, true, 0)

	boxTitulos.PackStart(labelGestacao, true, true, 0)
	boxEntradas.PackStart(textoGestacao, true, true, 0)

	boxTitulos.PackStart(labelGasCarbonico, true, true, 0)
	boxEntradas.PackStart(textoGasCarbonico, true, true, 0)

	boxTitulos.PackStart(labelGasOxigenio, true, true, 0)
	boxEntradas.PackStart(textoGasOxigenio, true, true, 0)

	btn := setup_btn("Salvar", func() {


		var nome, _ =	textoNome.GetText()
		var adulto ,_ =	textoAdulto.GetText()
		var vida,_ =	textoVida.GetText()
		var cor,_ =	textoCor.GetText()
		var freqreproducao,_ =	textoFreReproducao.GetText()
		var gestacao,_ =	textoGestacao.GetText()
		var gascarbonico,_ =	textoGasCarbonico.GetText()
		var gasoxigenio,_ =	textoGasOxigenio.GetText()

		var iadulto,_ = strconv.Atoi(adulto)
		var ivida,_ = strconv.Atoi(vida)
		var ifreqreproducao,_ = strconv.Atoi(freqreproducao)
		var icor,_ = strconv.Atoi(cor)
		var igestacao,_ = strconv.Atoi(gestacao)
		var igascarbonico,_ = strconv.Atoi(gascarbonico)
		var igasoxigenio,_ = strconv.Atoi(gasoxigenio)

		var tg=uint32(icor)


		x := &Organismo{
			Base: Base{Tipo: "Consumidor", Adulto: iadulto, Vida: ivida,Cor:tg},
			Reproducao:Reproducao{Frequencia:ifreqreproducao,Gestacao:igestacao},
			Taxas:Taxas{GasCarbonico:igascarbonico,GasOxigenio:igasoxigenio},

		}

		file, _ := os.Create( caminho + nome+".organismo" )
		xmlWriter := io.Writer(file)
		enc := xml.NewEncoder(xmlWriter)
		enc.Indent("", "  ")
		if err := enc.Encode(x); err != nil {
			fmt.Printf("error: %v\n", err)
		}


	})

	boxSalvadorProdutor.PackStart(btn, true, true, 0)



}

