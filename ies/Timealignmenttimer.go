package ies

import "rrc/utils"

// TimeAlignmentTimer ::= ENUMERATED
type Timealignmenttimer struct {
	Value utils.ENUMERATED
}

const (
	TimealignmenttimerEnumeratedNothing = iota
	TimealignmenttimerEnumeratedSf500
	TimealignmenttimerEnumeratedSf750
	TimealignmenttimerEnumeratedSf1280
	TimealignmenttimerEnumeratedSf1920
	TimealignmenttimerEnumeratedSf2560
	TimealignmenttimerEnumeratedSf5120
	TimealignmenttimerEnumeratedSf10240
	TimealignmenttimerEnumeratedInfinity
)
