package ies

import "rrc/utils"

// PDSCH-ConfigCommon-v1310-pdsch-maxNumRepetitionCEmodeA-r13 ::= ENUMERATED
type PdschConfigcommonV1310PdschMaxnumrepetitioncemodeaR13 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigcommonV1310PdschMaxnumrepetitioncemodeaR13EnumeratedNothing = iota
	PdschConfigcommonV1310PdschMaxnumrepetitioncemodeaR13EnumeratedR16
	PdschConfigcommonV1310PdschMaxnumrepetitioncemodeaR13EnumeratedR32
)
