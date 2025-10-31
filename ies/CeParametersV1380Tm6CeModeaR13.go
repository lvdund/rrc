package ies

import "rrc/utils"

// CE-Parameters-v1380-tm6-CE-ModeA-r13 ::= ENUMERATED
type CeParametersV1380Tm6CeModeaR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1380Tm6CeModeaR13EnumeratedNothing = iota
	CeParametersV1380Tm6CeModeaR13EnumeratedSupported
)
