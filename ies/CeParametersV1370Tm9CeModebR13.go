package ies

import "rrc/utils"

// CE-Parameters-v1370-tm9-CE-ModeB-r13 ::= ENUMERATED
type CeParametersV1370Tm9CeModebR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1370Tm9CeModebR13EnumeratedNothing = iota
	CeParametersV1370Tm9CeModebR13EnumeratedSupported
)
