package ies

import "rrc/utils"

// RRCConnectionResumeRequest-5GC-r15-IEs ::= SEQUENCE
type Rrcconnectionresumerequest5gcR15Ies struct {
	ResumeidentityR15  interface{}
	ShortresumemacIR15 utils.BITSTRING
	ResumecauseR15     ResumecauseR15
	Spare              utils.BITSTRING
}
