package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-multiRB-PUCCH-SCS-120kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17MultirbPucchScs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17MultirbPucchScs120khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17MultirbPucchScs120khzR17EnumeratedSupported
)
