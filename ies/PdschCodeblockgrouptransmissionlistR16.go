package ies

// PDSCH-CodeBlockGroupTransmissionList-r16 ::= SEQUENCE OF PDSCH-CodeBlockGroupTransmission
// SIZE (1..2)
type PdschCodeblockgrouptransmissionlistR16 struct {
	Value []PdschCodeblockgrouptransmission `lb:1,ub:2`
}
