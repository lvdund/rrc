package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-CP-EPC-r16 ::= ENUMERATED
type PurParametersNbR16PurCpEpcR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurCpEpcR16EnumeratedNothing = iota
	PurParametersNbR16PurCpEpcR16EnumeratedSupported
)
