package ies

import "rrc/utils"

// SPT-Parameters-r15 ::= SEQUENCE
type SptParametersR15 struct {
	FramestructuretypeSptR15 *utils.BITSTRING `lb:3,ub:3`
	MaxnumberccsSptR15       *utils.INTEGER   `lb:0,ub:32`
}
