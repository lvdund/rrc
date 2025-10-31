package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-widebandPRACH-SCS-120kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17WidebandprachScs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17WidebandprachScs120khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17WidebandprachScs120khzR17EnumeratedSupported
)
