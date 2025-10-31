package ies

import "rrc/utils"

// MAC-Parameters-v1630-directSCG-SCellActivationNEDC-r16 ::= ENUMERATED
type MacParametersV1630DirectscgScellactivationnedcR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1630DirectscgScellactivationnedcR16EnumeratedNothing = iota
	MacParametersV1630DirectscgScellactivationnedcR16EnumeratedSupported
)
