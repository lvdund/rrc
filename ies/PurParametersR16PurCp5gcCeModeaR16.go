package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-CP-5GC-CE-ModeA-r16 ::= ENUMERATED
type PurParametersR16PurCp5gcCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurCp5gcCeModeaR16EnumeratedNothing = iota
	PurParametersR16PurCp5gcCeModeaR16EnumeratedSupported
)
