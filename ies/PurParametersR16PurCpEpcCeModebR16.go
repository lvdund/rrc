package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-CP-EPC-CE-ModeB-r16 ::= ENUMERATED
type PurParametersR16PurCpEpcCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurCpEpcCeModebR16EnumeratedNothing = iota
	PurParametersR16PurCpEpcCeModebR16EnumeratedSupported
)
