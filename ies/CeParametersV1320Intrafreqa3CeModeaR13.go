package ies

import "rrc/utils"

// CE-Parameters-v1320-intraFreqA3-CE-ModeA-r13 ::= ENUMERATED
type CeParametersV1320Intrafreqa3CeModeaR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1320Intrafreqa3CeModeaR13EnumeratedNothing = iota
	CeParametersV1320Intrafreqa3CeModeaR13EnumeratedSupported
)
