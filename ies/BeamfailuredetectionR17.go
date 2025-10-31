package ies

// BeamFailureDetection-r17 ::= SEQUENCE
type BeamfailuredetectionR17 struct {
	Failuredetectionset1R17 *BeamfailuredetectionsetR17
	Failuredetectionset2R17 *BeamfailuredetectionsetR17
	AdditionalpciR17        *AdditionalpciindexR17
}
