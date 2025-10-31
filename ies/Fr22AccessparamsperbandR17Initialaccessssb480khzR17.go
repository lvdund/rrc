package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-initialAccessSSB-480kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17Initialaccessssb480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17Initialaccessssb480khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17Initialaccessssb480khzR17EnumeratedSupported
)
