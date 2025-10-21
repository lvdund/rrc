package ies

import "rrc/utils"

// TimeAlignmentTimer ::= ENUMERATED
type Timealignmenttimer struct {
	Value utils.ENUMERATED
}

const (
	TimealignmenttimerSf500    = 0
	TimealignmenttimerSf750    = 1
	TimealignmenttimerSf1280   = 2
	TimealignmenttimerSf1920   = 3
	TimealignmenttimerSf2560   = 4
	TimealignmenttimerSf5120   = 5
	TimealignmenttimerSf10240  = 6
	TimealignmenttimerInfinity = 7
)
