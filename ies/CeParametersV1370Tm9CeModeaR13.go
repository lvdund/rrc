package ies

import "rrc/utils"

// CE-Parameters-v1370-tm9-CE-ModeA-r13 ::= ENUMERATED
type CeParametersV1370Tm9CeModeaR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1370Tm9CeModeaR13EnumeratedNothing = iota
	CeParametersV1370Tm9CeModeaR13EnumeratedSupported
)
