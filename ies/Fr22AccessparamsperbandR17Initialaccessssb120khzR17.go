package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-initialAccessSSB-120kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17Initialaccessssb120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17Initialaccessssb120khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17Initialaccessssb120khzR17EnumeratedSupported
)
