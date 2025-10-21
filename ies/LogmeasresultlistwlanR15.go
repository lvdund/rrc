package ies

import "rrc/utils"

// LogMeasResultListWLAN-r15 ::= SEQUENCE OF LogMeasResultWLAN-r15
// SIZE (1..maxWLAN-Id-Report-r14)
type LogmeasresultlistwlanR15 struct {
	Value []LogmeasresultwlanR15
}
