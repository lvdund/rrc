package ies

// DCI7-CandidatesPerAL-SPDCCH-r15 ::= SEQUENCE OF DCI7-Candidates-r15
// SIZE (1..4)
type Dci7CandidatesperalSpdcchR15 struct {
	Value []Dci7CandidatesR15 `lb:1,ub:4`
}
