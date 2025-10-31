package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subcarrierPuncturingCE-ModeB-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubcarrierpuncturingceModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubcarrierpuncturingceModebR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubcarrierpuncturingceModebR16EnumeratedSupported
)
