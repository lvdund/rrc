package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-widebandPRACH-SCS-480kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17WidebandprachScs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17WidebandprachScs480khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17WidebandprachScs480khzR17EnumeratedSupported
)
