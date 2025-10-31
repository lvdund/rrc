package ies

// ResultsPerCSI-RS-Index ::= SEQUENCE
type ResultspercsiRsIndex struct {
	CsiRsIndex   CsiRsIndex
	CsiRsResults *Measquantityresults
}
