package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subframeResourceResvUL-CE-ModeB-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubframeresourceresvulCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubframeresourceresvulCeModebR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubframeresourceresvulCeModebR16EnumeratedSupported
)
