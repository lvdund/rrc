package ies

// EUTRA-Parameters ::= SEQUENCE
// Extensible
type EutraParameters struct {
	Supportedbandlisteutra []Freqbandindicatoreutra `lb:1,ub:maxBandsEUTRA`
	EutraParameterscommon  *EutraParameterscommon
	EutraParametersxddDiff *EutraParametersxddDiff
}
