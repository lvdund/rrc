package ies

import "rrc/utils"

// LAA-Parameters-v1430-twoStepSchedulingTimingInfo-r14 ::= ENUMERATED
type LaaParametersV1430TwostepschedulingtiminginfoR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430TwostepschedulingtiminginfoR14EnumeratedNothing = iota
	LaaParametersV1430TwostepschedulingtiminginfoR14EnumeratedNplus1
	LaaParametersV1430TwostepschedulingtiminginfoR14EnumeratedNplus2
	LaaParametersV1430TwostepschedulingtiminginfoR14EnumeratedNplus3
)
