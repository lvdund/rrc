package ies

import "rrc/utils"

// MAC-Parameters-v1530-extendedLCID-Duplication-r15 ::= ENUMERATED
type MacParametersV1530ExtendedlcidDuplicationR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530ExtendedlcidDuplicationR15EnumeratedNothing = iota
	MacParametersV1530ExtendedlcidDuplicationR15EnumeratedSupported
)
