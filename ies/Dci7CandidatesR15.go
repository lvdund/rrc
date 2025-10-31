package ies

import "rrc/utils"

// DCI7-Candidates-r15 ::= utils.INTEGER (0..6)
type Dci7CandidatesR15 struct {
	Value utils.INTEGER `lb:0,ub:6`
}
