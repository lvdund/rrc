package ies

// MeasResultIdleNR-r16-measResultServingCell-r16 ::= SEQUENCE
type MeasresultidlenrR16MeasresultservingcellR16 struct {
	RsrpResultR16        *RsrpRange
	RsrqResultR16        *RsrqRange
	ResultsssbIndexesR16 *ResultsperssbIndexlistR16
}
