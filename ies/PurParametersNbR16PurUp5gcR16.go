package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-UP-5GC-r16 ::= ENUMERATED
type PurParametersNbR16PurUp5gcR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurUp5gcR16EnumeratedNothing = iota
	PurParametersNbR16PurUp5gcR16EnumeratedSupported
)
