package ies

import "rrc/utils"

// CandidateBeamRS-r16 ::= SEQUENCE
type CandidatebeamrsR16 struct {
	CandidatebeamConfigR16 CandidatebeamConfigR16
	ServingCellId          utils.INTEGER `lb:0,ub:maxNrofServingCells`
}
