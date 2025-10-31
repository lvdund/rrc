package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-CP-EPC-CE-ModeA-r16 ::= ENUMERATED
type PurParametersR16PurCpEpcCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurCpEpcCeModeaR16EnumeratedNothing = iota
	PurParametersR16PurCpEpcCeModeaR16EnumeratedSupported
)
