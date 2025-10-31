package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-pdsch-MultiTB-CE-ModeB-r16 ::= ENUMERATED
type CeMultitbParametersR16PdschMultitbCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16PdschMultitbCeModebR16EnumeratedNothing = iota
	CeMultitbParametersR16PdschMultitbCeModebR16EnumeratedSupported
)
