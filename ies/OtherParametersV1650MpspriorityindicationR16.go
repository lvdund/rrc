package ies

import "rrc/utils"

// Other-Parameters-v1650-mpsPriorityIndication-r16 ::= ENUMERATED
type OtherParametersV1650MpspriorityindicationR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1650MpspriorityindicationR16EnumeratedNothing = iota
	OtherParametersV1650MpspriorityindicationR16EnumeratedSupported
)
