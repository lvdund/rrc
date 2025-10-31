package ies

import "rrc/utils"

// CG-CandidateInfo-r17 ::= SEQUENCE
type CgCandidateinfoR17 struct {
	CgCandidateinfoidR17 CgCandidateinfoidR17
	CandidatecgConfigR17 utils.OCTETSTRING
}
