package ies

import "rrc/utils"

// PUSCH-ConfigCommon-v1310-pusch-maxNumRepetitionCEmodeA-r13 ::= ENUMERATED
type PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13EnumeratedNothing = iota
	PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13EnumeratedR8
	PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13EnumeratedR16
	PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13EnumeratedR32
)
