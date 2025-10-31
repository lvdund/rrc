package ies

// TDD-UL-DL-ConfigCommon ::= SEQUENCE
// Extensible
type TddUlDlConfigcommon struct {
	Referencesubcarrierspacing Subcarrierspacing
	Pattern1                   TddUlDlPattern
	Pattern2                   *TddUlDlPattern
}
