package ies

import "rrc/utils"

// ServingCellConfigCommon-ssb-PositionsInBurst ::= CHOICE
const (
	ServingcellconfigcommonSsbPositionsinburstChoiceNothing = iota
	ServingcellconfigcommonSsbPositionsinburstChoiceShortbitmap
	ServingcellconfigcommonSsbPositionsinburstChoiceMediumbitmap
	ServingcellconfigcommonSsbPositionsinburstChoiceLongbitmap
)

type ServingcellconfigcommonSsbPositionsinburst struct {
	Choice       uint64
	Shortbitmap  *utils.BITSTRING `lb:4,ub:4`
	Mediumbitmap *utils.BITSTRING `lb:8,ub:8`
	Longbitmap   *utils.BITSTRING `lb:64,ub:64`
}
