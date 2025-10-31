package ies

import "rrc/utils"

// SL-Parameters-v1540-sl-RateMatchingTBSScaling-r15 ::= ENUMERATED
type SlParametersV1540SlRatematchingtbsscalingR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1540SlRatematchingtbsscalingR15EnumeratedNothing = iota
	SlParametersV1540SlRatematchingtbsscalingR15EnumeratedSupported
)
