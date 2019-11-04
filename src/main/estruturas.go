package main

type Organismo struct {
	Base Base `xml:"Base"`
	Reproducao Reproducao `xml:"Reproducao"`
	Taxas Taxas `xml:"Taxas"`

}

type Base struct {
	Tipo string `xml:"Tipo,attr"`
	Adulto        int `xml:"Adulto,attr"`
	Vida       int `xml:"Vida,attr"`
	Cor       uint32 `xml:"Cor,attr"`

}

type Reproducao struct {
	Frequencia        int `xml:"Frequencia,attr"`
	Gestacao        int `xml:"Gestacao,attr"`
}

type Taxas struct {
	GasCarbonico        int `xml:"GasCarbonico,attr"`
	GasOxigenio        int `xml:"GasOxigenio,attr"`
}