package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-pusch-MultiTB-CE-ModeA-r16 ::= ENUMERATED
type CeMultitbParametersR16PuschMultitbCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16PuschMultitbCeModeaR16EnumeratedNothing = iota
	CeMultitbParametersR16PuschMultitbCeModeaR16EnumeratedSupported
)
