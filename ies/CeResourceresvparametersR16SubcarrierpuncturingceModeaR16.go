package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subcarrierPuncturingCE-ModeA-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubcarrierpuncturingceModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubcarrierpuncturingceModeaR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubcarrierpuncturingceModeaR16EnumeratedSupported
)
