package ies

import "rrc/utils"

// MAC-Parameters-v1530-directSCellActivation-r15 ::= ENUMERATED
type MacParametersV1530DirectscellactivationR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530DirectscellactivationR15EnumeratedNothing = iota
	MacParametersV1530DirectscellactivationR15EnumeratedSupported
)
