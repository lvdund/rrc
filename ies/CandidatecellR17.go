package ies

import "rrc/utils"

// CandidateCell-r17 ::= SEQUENCE
type CandidatecellR17 struct {
	PhyscellidR17           Physcellid
	CondexecutioncondscgR17 *utils.OCTETSTRING
}
