package ies

import "rrc/utils"

// MAC-Parameters-v1530-sps-ServingCell-r15 ::= ENUMERATED
type MacParametersV1530SpsServingcellR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530SpsServingcellR15EnumeratedNothing = iota
	MacParametersV1530SpsServingcellR15EnumeratedSupported
)
