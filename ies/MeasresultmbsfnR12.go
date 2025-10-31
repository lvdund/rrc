package ies

// MeasResultMBSFN-r12 ::= SEQUENCE
// Extensible
type MeasresultmbsfnR12 struct {
	MbsfnAreaR12             MeasresultmbsfnR12MbsfnAreaR12
	RsrpresultmbsfnR12       RsrpRange
	RsrqresultmbsfnR12       MbsfnRsrqRangeR12
	SignallingblerResultR12  *BlerResultR12
	DatablerMchResultlistR12 *DatablerMchResultlistR12
}
