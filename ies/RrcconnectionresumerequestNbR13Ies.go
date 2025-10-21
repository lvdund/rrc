package ies

import "rrc/utils"

// RRCConnectionResumeRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumerequestNbR13Ies struct {
	ResumeidR13                  ResumeidentityR13
	ShortresumemacIR13           ShortmacI
	ResumecauseR13               EstablishmentcauseNbR13
	EarlycontentionresolutionR14 bool
	CqiNpdcchR14                 CqiNpdcchNbR14
	AnrInfoavailableR16          bool
	Spare                        utils.BITSTRING
}
