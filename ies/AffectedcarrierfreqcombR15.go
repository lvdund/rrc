package ies

import "rrc/utils"

// AffectedCarrierFreqComb-r15 ::= SEQUENCE OF MeasObjectId-r13
// SIZE (1..maxServCell-r13)
type AffectedcarrierfreqcombR15 struct {
	Value []MeasobjectidR13
}
