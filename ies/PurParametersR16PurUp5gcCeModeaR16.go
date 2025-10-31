package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-UP-5GC-CE-ModeA-r16 ::= ENUMERATED
type PurParametersR16PurUp5gcCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurUp5gcCeModeaR16EnumeratedNothing = iota
	PurParametersR16PurUp5gcCeModeaR16EnumeratedSupported
)
