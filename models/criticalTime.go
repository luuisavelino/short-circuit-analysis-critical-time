package models

type Matrix [][]string

var Zbus Matrix

var AllZbusBeforeFault map[string]Matrix

var AllZbusAfterFault map[string]Matrix

type Element struct {
	Id         int    `json:"id"`
	De         string `json:"de"`
	Para       string `json:"para"`
	Nome       string `json:"nome"`
	Z_positiva string `json:"z_positiva"`
	Z_zero     string `json:"z_zero"`
}

var Elements = make(map[string]map[string]Element)

type Posicao_zbus struct {
	Posicao int
	Tipo    string
}

var BarrasAdicionadasBefore = make(map[string]Posicao_zbus)

var BarrasAdicionadasAfter = make(map[string]Posicao_zbus)