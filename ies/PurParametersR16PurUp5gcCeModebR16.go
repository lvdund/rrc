package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-UP-5GC-CE-ModeB-r16 ::= ENUMERATED
type PurParametersR16PurUp5gcCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurUp5gcCeModebR16EnumeratedNothing = iota
	PurParametersR16PurUp5gcCeModebR16EnumeratedSupported
)
