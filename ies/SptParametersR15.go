package ies

import "rrc/utils"

// SPT-Parameters-r15 ::= SEQUENCE
type SptParametersR15 struct {
	FramestructuretypeSptR15 *utils.BITSTRING
	MaxnumberccsSptR15       *utils.INTEGER
}
