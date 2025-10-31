package ies

import "rrc/utils"

// RRCConnectionResumeRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumerequestNbR13 struct {
	ResumeidR13                  ResumeidentityR13
	ShortresumemacIR13           ShortmacI
	ResumecauseR13               EstablishmentcauseNbR13
	EarlycontentionresolutionR14 utils.BOOLEAN
	CqiNpdcchR14                 CqiNpdcchNbR14
	AnrInfoavailableR16          utils.BOOLEAN
	Spare                        utils.BITSTRING `lb:3,ub:3`
}
