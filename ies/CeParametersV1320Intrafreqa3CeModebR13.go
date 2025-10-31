package ies

import "rrc/utils"

// CE-Parameters-v1320-intraFreqA3-CE-ModeB-r13 ::= ENUMERATED
type CeParametersV1320Intrafreqa3CeModebR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1320Intrafreqa3CeModebR13EnumeratedNothing = iota
	CeParametersV1320Intrafreqa3CeModebR13EnumeratedSupported
)
