package ies

// CandidateBeamRSListExt-r16 ::= SEQUENCE OF PRACH-ResourceDedicatedBFR
// SIZE (1.. maxNrofCandidateBeamsExt-r16)
type CandidatebeamrslistextR16 struct {
	Value []PrachResourcededicatedbfr `lb:1,ub:maxNrofCandidateBeamsExtR16`
}
