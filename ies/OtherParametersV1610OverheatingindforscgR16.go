package ies

import "rrc/utils"

// Other-Parameters-v1610-overheatingIndForSCG-r16 ::= ENUMERATED
type OtherParametersV1610OverheatingindforscgR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610OverheatingindforscgR16EnumeratedNothing = iota
	OtherParametersV1610OverheatingindforscgR16EnumeratedSupported
)
