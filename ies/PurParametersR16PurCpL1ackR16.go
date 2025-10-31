package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-CP-L1Ack-r16 ::= ENUMERATED
type PurParametersR16PurCpL1ackR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurCpL1ackR16EnumeratedNothing = iota
	PurParametersR16PurCpL1ackR16EnumeratedSupported
)
