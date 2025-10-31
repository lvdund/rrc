package ies

import "rrc/utils"

// RRCConnectionResumeRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcconnectionresumerequest5gcNbR16 struct {
	ResumeidR16        IRntiR15
	ShortresumemacIR16 ShortmacI
	ResumecauseR16     EstablishmentcauseNbR13
	CqiNpdcchR16       CqiNpdcchNbR14
	Spare              utils.BITSTRING `lb:4,ub:4`
}
