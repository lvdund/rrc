package ies

import "rrc/utils"

// P0AlphaSet-r17 ::= SEQUENCE
type P0alphasetR17 struct {
	P0R17              *utils.INTEGER `lb:0,ub:15`
	AlphaR17           *Alpha
	ClosedloopindexR17 P0alphasetR17ClosedloopindexR17
}
