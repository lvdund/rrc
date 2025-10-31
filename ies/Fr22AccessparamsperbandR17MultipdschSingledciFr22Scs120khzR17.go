package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-multiPDSCH-SingleDCI-FR2-2-SCS-120kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17MultipdschSingledciFr22Scs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17MultipdschSingledciFr22Scs120khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17MultipdschSingledciFr22Scs120khzR17EnumeratedSupported
)
