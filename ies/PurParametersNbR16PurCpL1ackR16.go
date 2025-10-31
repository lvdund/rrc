package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-CP-L1Ack-r16 ::= ENUMERATED
type PurParametersNbR16PurCpL1ackR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurCpL1ackR16EnumeratedNothing = iota
	PurParametersNbR16PurCpL1ackR16EnumeratedSupported
)
