package ies

// BeamFailureRecoveryRSConfig-r16 ::= SEQUENCE
// Extensible
type BeamfailurerecoveryrsconfigR16 struct {
	RsrpThresholdbfrR16     *RsrpRange
	CandidatebeamrsListR16  *[]CandidatebeamrsR16 `lb:1,ub:maxNrofCandidateBeamsR16`
	CandidatebeamrsList2R17 *[]CandidatebeamrsR16 `lb:1,ub:maxNrofCandidateBeamsR16,ext`
}
