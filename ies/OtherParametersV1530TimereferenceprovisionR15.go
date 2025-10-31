package ies

import "rrc/utils"

// Other-Parameters-v1530-timeReferenceProvision-r15 ::= ENUMERATED
type OtherParametersV1530TimereferenceprovisionR15 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1530TimereferenceprovisionR15EnumeratedNothing = iota
	OtherParametersV1530TimereferenceprovisionR15EnumeratedSupported
)
