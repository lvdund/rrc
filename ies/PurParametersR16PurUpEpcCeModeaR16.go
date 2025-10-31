package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-UP-EPC-CE-ModeA-r16 ::= ENUMERATED
type PurParametersR16PurUpEpcCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurUpEpcCeModeaR16EnumeratedNothing = iota
	PurParametersR16PurUpEpcCeModeaR16EnumeratedSupported
)
