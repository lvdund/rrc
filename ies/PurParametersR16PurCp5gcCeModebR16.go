package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-CP-5GC-CE-ModeB-r16 ::= ENUMERATED
type PurParametersR16PurCp5gcCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurCp5gcCeModebR16EnumeratedNothing = iota
	PurParametersR16PurCp5gcCeModebR16EnumeratedSupported
)
