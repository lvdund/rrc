package ies

import "rrc/utils"

// DummyB ::= SEQUENCE
type Dummyb struct {
	Maxnumbertxportsperresource  DummybMaxnumbertxportsperresource
	Maxnumberresources           utils.INTEGER `lb:0,ub:64`
	Totalnumbertxports           utils.INTEGER `lb:0,ub:256`
	Supportedcodebookmode        DummybSupportedcodebookmode
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
