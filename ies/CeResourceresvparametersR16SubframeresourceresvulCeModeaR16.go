package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-subframeResourceResvUL-CE-ModeA-r16 ::= ENUMERATED
type CeResourceresvparametersR16SubframeresourceresvulCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SubframeresourceresvulCeModeaR16EnumeratedNothing = iota
	CeResourceresvparametersR16SubframeresourceresvulCeModeaR16EnumeratedSupported
)
