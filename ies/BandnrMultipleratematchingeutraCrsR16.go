package ies

import "rrc/utils"

// BandNR-multipleRateMatchingEUTRA-CRS-r16 ::= SEQUENCE
type BandnrMultipleratematchingeutraCrsR16 struct {
	MaxnumberpatternsR16           utils.INTEGER `lb:0,ub:6`
	MaxnumbernonOverlappatternsR16 utils.INTEGER `lb:0,ub:3`
}
