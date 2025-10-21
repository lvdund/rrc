package ies

import "rrc/utils"

// SL-ReliabilityList-r15 ::= SEQUENCE OF SL-Reliability-r15
// SIZE (1..maxSL-Reliability-r15)
type SlReliabilitylistR15 struct {
	Value []SlReliabilityR15
}
