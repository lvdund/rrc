package ies

import "rrc/utils"

// DummyC ::= SEQUENCE
type Dummyc struct {
	Maxnumbertxportsperresource  DummycMaxnumbertxportsperresource
	Maxnumberresources           utils.INTEGER `lb:0,ub:64`
	Totalnumbertxports           utils.INTEGER `lb:0,ub:256`
	Supportedcodebookmode        DummycSupportedcodebookmode
	Supportednumberpanels        DummycSupportednumberpanels
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
