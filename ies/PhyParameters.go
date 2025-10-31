package ies

// Phy-Parameters ::= SEQUENCE
type PhyParameters struct {
	PhyParameterscommon  *PhyParameterscommon
	PhyParametersxddDiff *PhyParametersxddDiff
	PhyParametersfrxDiff *PhyParametersfrxDiff
	PhyParametersfr1     *PhyParametersfr1
	PhyParametersfr2     *PhyParametersfr2
}
