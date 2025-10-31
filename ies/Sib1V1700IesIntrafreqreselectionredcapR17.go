package ies

import "rrc/utils"

// SIB1-v1700-IEs-intraFreqReselectionRedCap-r17 ::= ENUMERATED
type Sib1V1700IesIntrafreqreselectionredcapR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib1V1700IesIntrafreqreselectionredcapR17EnumeratedNothing = iota
	Sib1V1700IesIntrafreqreselectionredcapR17EnumeratedAllowed
	Sib1V1700IesIntrafreqreselectionredcapR17EnumeratedNotallowed
)
