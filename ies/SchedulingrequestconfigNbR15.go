package ies

import "rrc/utils"

// SchedulingRequestConfig-NB-r15 ::= SEQUENCE
// Extensible
type SchedulingrequestconfigNbR15 struct {
	SrWithharqAckConfigR15    *utils.ENUMERATED
	SrWithoutharqAckConfigR15 *SrWithoutharqAckConfigNbR15
	SrSpsBsrConfigR15         *SrSpsBsrConfigNbR15
}
