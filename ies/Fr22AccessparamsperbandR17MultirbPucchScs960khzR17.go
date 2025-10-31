package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-multiRB-PUCCH-SCS-960kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17MultirbPucchScs960khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17MultirbPucchScs960khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17MultirbPucchScs960khzR17EnumeratedSupported
)
