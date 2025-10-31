package ies

// MeasResultsPerCellIdleNR-r16-measIdleResultNR-r16 ::= SEQUENCE
type MeasresultspercellidlenrR16MeasidleresultnrR16 struct {
	RsrpResultR16        *RsrpRange
	RsrqResultR16        *RsrqRange
	ResultsssbIndexesR16 *ResultsperssbIndexlistR16
}
