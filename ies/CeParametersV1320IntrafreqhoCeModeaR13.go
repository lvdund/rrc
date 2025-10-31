package ies

import "rrc/utils"

// CE-Parameters-v1320-intraFreqHO-CE-ModeA-r13 ::= ENUMERATED
type CeParametersV1320IntrafreqhoCeModeaR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1320IntrafreqhoCeModeaR13EnumeratedNothing = iota
	CeParametersV1320IntrafreqhoCeModeaR13EnumeratedSupported
)
