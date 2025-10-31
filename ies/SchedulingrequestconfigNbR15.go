package ies

// SchedulingRequestConfig-NB-r15 ::= SEQUENCE
// Extensible
type SchedulingrequestconfigNbR15 struct {
	SrWithharqAckConfigR15    *bool
	SrWithoutharqAckConfigR15 *SrWithoutharqAckConfigNbR15
	SrSpsBsrConfigR15         *SrSpsBsrConfigNbR15
}
