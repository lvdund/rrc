package ies

import "rrc/utils"

// DummyD ::= SEQUENCE
type Dummyd struct {
	Maxnumbertxportsperresource  DummydMaxnumbertxportsperresource
	Maxnumberresources           utils.INTEGER `lb:0,ub:64`
	Totalnumbertxports           utils.INTEGER `lb:0,ub:256`
	Parameterlx                  utils.INTEGER `lb:0,ub:4`
	Amplitudescalingtype         DummydAmplitudescalingtype
	Amplitudesubsetrestriction   *DummydAmplitudesubsetrestriction
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
