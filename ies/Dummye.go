package ies

import "rrc/utils"

// DummyE ::= SEQUENCE
type Dummye struct {
	Maxnumbertxportsperresource  DummyeMaxnumbertxportsperresource
	Maxnumberresources           utils.INTEGER `lb:0,ub:64`
	Totalnumbertxports           utils.INTEGER `lb:0,ub:256`
	Parameterlx                  utils.INTEGER `lb:0,ub:4`
	Amplitudescalingtype         DummyeAmplitudescalingtype
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
