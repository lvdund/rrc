package ies

import "rrc/utils"

// TimeAlignmentTimer ::= ENUMERATED
type Timealignmenttimer struct {
	Value utils.ENUMERATED
}

const (
	TimealignmenttimerEnumeratedNothing = iota
	TimealignmenttimerEnumeratedMs500
	TimealignmenttimerEnumeratedMs750
	TimealignmenttimerEnumeratedMs1280
	TimealignmenttimerEnumeratedMs1920
	TimealignmenttimerEnumeratedMs2560
	TimealignmenttimerEnumeratedMs5120
	TimealignmenttimerEnumeratedMs10240
	TimealignmenttimerEnumeratedInfinity
)
