package ies

// Phy-ParametersXDD-Diff ::= SEQUENCE
// Extensible
type PhyParametersxddDiff struct {
	Dynamicsfi                   *PhyParametersxddDiffDynamicsfi
	TwopucchF02Consecsymbols     *PhyParametersxddDiffTwopucchF02Consecsymbols
	TwodifferenttpcLoopPusch     *PhyParametersxddDiffTwodifferenttpcLoopPusch
	TwodifferenttpcLoopPucch     *PhyParametersxddDiffTwodifferenttpcLoopPucch
	DlSchedulingoffsetPdschTypea *PhyParametersxddDiffDlSchedulingoffsetPdschTypea `ext`
	DlSchedulingoffsetPdschTypeb *PhyParametersxddDiffDlSchedulingoffsetPdschTypeb `ext`
	UlSchedulingoffset           *PhyParametersxddDiffUlSchedulingoffset           `ext`
}
