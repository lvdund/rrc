package ies

import "rrc/utils"

// RRCConnectionResumeRequest-r13-IEs ::= SEQUENCE
type RrcconnectionresumerequestR13Ies struct {
	ResumeidentityR13  interface{}
	ShortresumemacIR13 utils.BITSTRING
	ResumecauseR13     Resumecause
	Spare              utils.BITSTRING
}
