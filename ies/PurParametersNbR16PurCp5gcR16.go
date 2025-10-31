package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-CP-5GC-r16 ::= ENUMERATED
type PurParametersNbR16PurCp5gcR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurCp5gcR16EnumeratedNothing = iota
	PurParametersNbR16PurCp5gcR16EnumeratedSupported
)
