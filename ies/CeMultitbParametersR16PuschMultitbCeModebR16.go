package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-pusch-MultiTB-CE-ModeB-r16 ::= ENUMERATED
type CeMultitbParametersR16PuschMultitbCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16PuschMultitbCeModebR16EnumeratedNothing = iota
	CeMultitbParametersR16PuschMultitbCeModebR16EnumeratedSupported
)
