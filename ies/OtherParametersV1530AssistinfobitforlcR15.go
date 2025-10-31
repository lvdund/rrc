package ies

import "rrc/utils"

// Other-Parameters-v1530-assistInfoBitForLC-r15 ::= ENUMERATED
type OtherParametersV1530AssistinfobitforlcR15 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1530AssistinfobitforlcR15EnumeratedNothing = iota
	OtherParametersV1530AssistinfobitforlcR15EnumeratedSupported
)
