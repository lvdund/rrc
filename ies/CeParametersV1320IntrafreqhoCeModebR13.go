package ies

import "rrc/utils"

// CE-Parameters-v1320-intraFreqHO-CE-ModeB-r13 ::= ENUMERATED
type CeParametersV1320IntrafreqhoCeModebR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1320IntrafreqhoCeModebR13EnumeratedNothing = iota
	CeParametersV1320IntrafreqhoCeModebR13EnumeratedSupported
)
