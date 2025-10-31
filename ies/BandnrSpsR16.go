package ies

import "rrc/utils"

// BandNR-sps-r16 ::= SEQUENCE
type BandnrSpsR16 struct {
	MaxnumberconfigsperbwpR16 utils.INTEGER `lb:0,ub:8`
	MaxnumberconfigsallccR16  utils.INTEGER `lb:0,ub:32`
}
