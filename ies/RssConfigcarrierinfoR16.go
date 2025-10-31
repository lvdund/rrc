package ies

import "rrc/utils"

// RSS-ConfigCarrierInfo-r16 ::= SEQUENCE
type RssConfigcarrierinfoR16 struct {
	NarrowbandindexR16       utils.BITSTRING `lb:1,ub:maxAvailNarrowBands1R16`
	TimeoffsetgranularityR16 RssConfigcarrierinfoR16TimeoffsetgranularityR16
}
