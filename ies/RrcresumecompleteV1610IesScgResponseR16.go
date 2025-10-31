package ies

import "rrc/utils"

// RRCResumeComplete-v1610-IEs-scg-Response-r16 ::= CHOICE
const (
	RrcresumecompleteV1610IesScgResponseR16ChoiceNothing = iota
	RrcresumecompleteV1610IesScgResponseR16ChoiceNrScgResponse
	RrcresumecompleteV1610IesScgResponseR16ChoiceEutraScgResponse
)

type RrcresumecompleteV1610IesScgResponseR16 struct {
	Choice           uint64
	NrScgResponse    *utils.OCTETSTRING
	EutraScgResponse *utils.OCTETSTRING
}
