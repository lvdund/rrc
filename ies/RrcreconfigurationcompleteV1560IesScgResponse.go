package ies

import "rrc/utils"

// RRCReconfigurationComplete-v1560-IEs-scg-Response ::= CHOICE
const (
	RrcreconfigurationcompleteV1560IesScgResponseChoiceNothing = iota
	RrcreconfigurationcompleteV1560IesScgResponseChoiceNrScgResponse
	RrcreconfigurationcompleteV1560IesScgResponseChoiceEutraScgResponse
)

type RrcreconfigurationcompleteV1560IesScgResponse struct {
	Choice           uint64
	NrScgResponse    *utils.OCTETSTRING
	EutraScgResponse *utils.OCTETSTRING
}
