package ies

import "rrc/utils"

// Other-Parameters-v1430-bwPrefInd-r14 ::= ENUMERATED
type OtherParametersV1430BwprefindR14 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1430BwprefindR14EnumeratedNothing = iota
	OtherParametersV1430BwprefindR14EnumeratedSupported
)
