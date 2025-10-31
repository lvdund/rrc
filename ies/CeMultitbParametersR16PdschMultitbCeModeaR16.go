package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-pdsch-MultiTB-CE-ModeA-r16 ::= ENUMERATED
type CeMultitbParametersR16PdschMultitbCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16PdschMultitbCeModeaR16EnumeratedNothing = iota
	CeMultitbParametersR16PdschMultitbCeModeaR16EnumeratedSupported
)
