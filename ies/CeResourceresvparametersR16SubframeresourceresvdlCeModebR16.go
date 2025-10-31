package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subframeResourceResvDL-CE-ModeB-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubframeresourceresvdlCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubframeresourceresvdlCeModebR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubframeresourceresvdlCeModebR16EnumeratedSupported
)
