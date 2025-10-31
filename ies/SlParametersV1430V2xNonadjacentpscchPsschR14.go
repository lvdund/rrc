package ies

import "rrc/utils"

// SL-Parameters-v1430-v2x-nonAdjacentPSCCH-PSSCH-r14 ::= ENUMERATED
type SlParametersV1430V2xNonadjacentpscchPsschR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430V2xNonadjacentpscchPsschR14EnumeratedNothing = iota
	SlParametersV1430V2xNonadjacentpscchPsschR14EnumeratedSupported
)
