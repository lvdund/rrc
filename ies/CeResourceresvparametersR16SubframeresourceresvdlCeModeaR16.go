package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subframeResourceResvDL-CE-ModeA-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubframeresourceresvdlCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubframeresourceresvdlCeModeaR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubframeresourceresvdlCeModeaR16EnumeratedSupported
)
