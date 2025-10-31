package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-SubPRB-CE-ModeA-r16 ::= ENUMERATED
type PurParametersR16PurSubprbCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurSubprbCeModeaR16EnumeratedNothing = iota
	PurParametersR16PurSubprbCeModeaR16EnumeratedSupported
)
