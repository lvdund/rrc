package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-UP-EPC-r16 ::= ENUMERATED
type PurParametersNbR16PurUpEpcR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurUpEpcR16EnumeratedNothing = iota
	PurParametersNbR16PurUpEpcR16EnumeratedSupported
)
