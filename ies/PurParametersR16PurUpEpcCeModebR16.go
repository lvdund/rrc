package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-UP-EPC-CE-ModeB-r16 ::= ENUMERATED
type PurParametersR16PurUpEpcCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurUpEpcCeModebR16EnumeratedNothing = iota
	PurParametersR16PurUpEpcCeModebR16EnumeratedSupported
)
