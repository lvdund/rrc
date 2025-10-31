package ies

import "rrc/utils"

// GWUS-GroupNarrowBandList-r16 ::= SEQUENCE OF utils.BOOLEAN // SIZE (1..maxAvailNarrowBands-r13)
type GwusGroupnarrowbandlistR16 struct {
	Value []utils.BOOLEAN `lb:1,ub:maxAvailNarrowBandsR13`
}
