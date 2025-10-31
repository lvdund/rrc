package ies

// ResultsPerSSB-Index ::= SEQUENCE
type ResultsperssbIndex struct {
	SsbIndex   SsbIndex
	SsbResults *Measquantityresults
}
